package main

import (
	"log"

	pb "github.com/ercsnmrs/grpc_reference/calculator/proto"
)

// Search for CaclulatorServiceServer
func (s *Server) Primes(req *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", req.Number)

	divisor := int32(2)
	number := req.Number

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: divisor,
			})
			number /= divisor
			log.Printf("Number Value: %v\n", number)
		} else {
			divisor++
		}
	}

	return nil
}
