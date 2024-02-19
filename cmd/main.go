package main

import (
	"github.com/Terralayr/platform/internal/service"
	"log"
)

func main() {
	svc, err := service.New()
	if err != nil {
		log.Println(err)
		return
	}

	user, err := svc.CreateUser()
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return
	}
	log.Printf("User: %v", user)

	physicalAsset, err := svc.CreatePhysicalAsset()
	if err != nil {
		log.Printf("Error creating physical asset: %v", err)
		return
	}
	log.Printf("Physical asset: %v", physicalAsset)
	log.Printf("Physical asset: %v", physicalAsset.String())
	log.Printf("Physical asset: %v", physicalAsset.Name())
	log.Printf("Physical asset: %v", physicalAsset.EnergyCapacity())
	log.Printf("Physical asset: %v", physicalAsset.LiveDispatch())

	block, err := svc.CreateBlock(
		physicalAsset.ID(),
		user.ID,
	)
	if err != nil {
		log.Printf("Error creating block: %v", err)
		return
	}
	log.Printf("Block: %v", block)

	log.Println("Setup complete")
}
