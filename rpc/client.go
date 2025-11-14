package rpc

import "google.golang.org/grpc"

func NewClient(address string) (*grpc.ClientConn, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
