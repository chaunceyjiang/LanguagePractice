package main
/*
import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:37017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	var db *mongo.Database
	db = client.Database("cbms")
	coll := db.Collection("main_app_datapath")
	cur := coll.FindOne(ctx, bson.M{"state": "apply"})
	var result Topov
	if err = cur.Decode(&result);err!=nil{
		log.Error(err)
	}
	fmt.Println(result.GetTwoWayTrandeIntfIPS("intf1"))
	coll = db.Collection("instance_id")
	curs,err :=coll.Find(ctx,bson.D{})
	if err!=nil{
		log.Error(err)
	}
	var instanceId bson.M
	for curs.Next(ctx){
		if err = curs.Decode(&instanceId);err!=nil{
			log.Error(err)
		}
		fmt.Println(instanceId)
	}
}*/
