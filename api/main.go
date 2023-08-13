package main

import (
	"canaanadvisors-test/core/repositories"
	"canaanadvisors-test/proto/management"
	"canaanadvisors-test/proto/notification"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"os/signal"


	"canaanadvisors-test/config"
	"canaanadvisors-test/core/app"
	"canaanadvisors-test/infra"
	"canaanadvisors-test/proto/order"
	"canaanadvisors-test/proto/user"
)

func main() {
	config.ReadConfig()
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)
	go func() {
		// Wait for termination signal
		<-signalChan
		// Trigger cancellation of the context
		cancel()
		// Wait for goroutine to finish
		fmt.Println("The service terminated gracefully")
	}()


	ap := fx.New(
		fx.Provide(
			infra.NewLogger,
			infra.NewTemporalClient,
			infra.NewDB,
			context.TODO,
			// TODO add all providers
			app.NewOrder,
			app.NewUser,
			app.NewNotification,
			app.NewManagement,
			repositories.NewOrderRepository,
			config.LoadTempoConfig,
		),
		fx.Invoke(
			listenAndServe,
		),
	)
	if err := ap.Start(ctx); err != nil {
		os.Exit(1)
	}
}

func listenAndServe(ctx context.Context, logger *zap.Logger,
	orderApp app.Order, authApp app.User, mgtApp app.Management, notifyApp app.Notification) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.C.Server.GRPCPort))
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Create a gRPC server instance
	grpcServer := grpc.NewServer()
	// Register our service with the gRPC server
	order.RegisterOrderServiceServer(grpcServer, &OrderController{logger: logger, app: orderApp})
	user.RegisterUserServiceServer(grpcServer, &UserController{logger: logger, app: authApp})
	notification.RegisterNotificationServiceServer(grpcServer, &NotificationController{logger: logger, app: notifyApp})
	notification.RegisterWebSocketServiceServer(grpcServer, &WebSocketController{logger: logger})
	management.RegisterManagementServiceServer(grpcServer, &ManagementController{logger: logger, app: mgtApp})
	// TODO: register more service here

	// Serve gRPC server
	logger.Info(fmt.Sprintf("Serving gRPC on 0.0.0.0:%v", config.C.Server.GRPCPort))
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	maxMsgSize := 1024 * 1024 * 20
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("0.0.0.0:%v", config.C.Server.GRPCPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	gwMux := runtime.NewServeMux()
	// Register service handlers
	err = order.RegisterOrderServiceHandler(ctx, gwMux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = user.RegisterUserServiceHandler(ctx, gwMux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = notification.RegisterNotificationServiceHandler(ctx, gwMux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = management.RegisterManagementServiceHandler(ctx, gwMux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = notification.RegisterWebSocketServiceHandler(ctx, gwMux, conn)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// TODO: register more service handlers here

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.C.Server.HTTPPort),
		Handler: gwMux,
	}

	logger.Info(fmt.Sprintf("Serving gRPC-Gateway on port %v", config.C.Server.HTTPPort))
	go func() {
		if err = gwServer.ListenAndServe(); err != nil {
			logger.Fatal(err.Error())
		}
	}()
	// Wait for a signal to shut down the server
	<-ctx.Done()

	// Gracefully stop the server
	grpcServer.GracefulStop()
}

