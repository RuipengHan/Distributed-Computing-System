package client_model

import (
	"context"
	"log"

	"cs425mp4/grpcclient/fileclient"
	"cs425mp4/grpcclient/pythonclient"
	"cs425mp4/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client asks Server to setup the model in all virtual machines
func ClientStartJob(addr string, job_id int, batch_size int, model_type string) (string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("ClientStartJob() did not connect: %v", err)
		return "", err
	}
	defer conn.Close()

	client := fileclient.NewClient(conn, nil)
	res_status, err := client.StartJob(context.Background(), job_id, batch_size, model_type)
	if res_status != "OK" {
		log.Print("Failed to initialize model")
	}
	return res_status, err
}

// Client asks Server to start inferencing on Model with id = Job Id
func ClientInferenceJob(addr string, job_id int) (string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("ClientInferenceJob() did not connect: %v", err)
		return "", err
	}
	defer conn.Close()

	client := fileclient.NewClient(conn, nil)
	res_status, err := client.StartInference(context.Background(), job_id)
	if res_status != "OK" {
		log.Print("Failed to initialize model")
	}
	return res_status, err
}

// Client asks Server to remove Model with id = Job Id
func ClientRemoveModel(addr string, job_id int) (string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("ClientRemoveModel() did not connect: %v", err)
		return "", err
	}
	defer conn.Close()

	client := fileclient.NewClient(conn, nil)
	res_status, err := client.RequestRemove(context.Background(), job_id)
	if res_status != "OK" {
		log.Print("Failed to remove model")
	}
	return res_status, err
}

// Server send job information for member to request files to download
func SendInferenceInformation(addr string, job_id int, batch_id int, file_replicas map[string][]string) string {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("SendInferenceInformation() did not connect: %v", err)
		return "Connection Failed"
	}
	defer conn.Close()
	client := fileclient.NewClient(conn, nil)
	res_status, err := client.SendJobInformation(context.Background(), batch_id, job_id, file_replicas)
	if res_status != "OK" {
		log.Printf("SendInferenceInformation() failed")
		return "Send Information Failed Failed"
	}
	log.Print("Good status")
	return "OK"
}

// Master asks Members in the Membership list to initialize their models
func AskMemberToInitializeModels(addr string, job_id int, batch_size int, model_type string) (string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("SendInferenceInformation() did not connect: %v", err)
		return "Connection Failed", err
	}
	defer conn.Close()
	client := fileclient.NewClient(conn, nil)
	res, err := client.AskMemberToInitializeModels(context.Background(), job_id, batch_size, model_type)
	if err != nil || res != "OK" {
		log.Printf("AskMemberToInitializeModels() failed")
		return "Send Information Failed Failed", err
	}
	return "OK", nil
}

func SendJobStatusReplication(addr string, job_status utils.JobStatus) (string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("SendInferenceInformation() did not connect: %v", err)
		return "Connection Failed", err
	}
	defer conn.Close()
	client := fileclient.NewClient(conn, nil)
	res, err := client.SendJobStatusReplication(context.Background(), job_status)
	if err != nil || res != "OK" {
		log.Printf("SendJobStatusReplication() failed")
		return "Failed to send job status replication", err
	}
	return res, err
}

// Go Member ask Python Server to initialize model
func AskToInitializeModel(addr string, job_id int, batch_size int, model_type string) (string, error) {
	log.Println("Requested to initialize model job")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("ClientAskToReplicate() did not connect: %v", err)
		return "", err
	}
	defer conn.Close()
	client := pythonclient.NewPythonClient(conn)
	res_status, err := client.InitializeModel(context.Background(), job_id, batch_size, model_type)
	if res_status != "OK" {
		log.Print("Failed to initialize model")
	}
	return res_status, err
}

// Go Member Ask Python Server to inference on takss
func AskToInference(addr string, job_id int, batch_id int, inference_size int, data_folder string) ([]byte, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("AskToInference() did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pythonclient.NewPythonClient(conn)
	res_status, err := client.ModelInference(context.Background(), job_id, batch_id, inference_size, data_folder)
	return res_status, nil
}

// Server ask each member to remove models
func AskMemberToRemoveModels(addr string, job_id int) (string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("AskToInference() did not connect: %v", err)
		return "", err
	}
	defer conn.Close()
	client := fileclient.NewClient(conn, nil)
	res, err := client.AskMemberToRemoveModels(context.Background(), job_id)
	return res, nil
}

// Go Member ask PythonS erver to remove models
func AskToRemoveModels(addr string, job_id int) (string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("AskToRemoveModels() did not connect: %v", err)
		return "", err
	}
	defer conn.Close()
	client := pythonclient.NewPythonClient(conn)
	res, err := client.RemoveModel(context.Background(), job_id)
	return res, nil
}

// Client request to get all informations
func ClientRequestJobStatus(addr string, job_id int) (string, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("ClientRequestJobStatus() did not connect: %v", err)
		return "", err
	}
	defer conn.Close()
	client := fileclient.NewClient(conn, nil)
	res, err := client.RequestJobStatus(context.Background(), job_id)
	return res, nil
}
