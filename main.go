package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----------------------------------------")
	fmt.Println("        Trade Engine 1.0               ")
	fmt.Println("-----------------------------------------")

	// create consumer for consuming order
	// consumer := createConsumer()

	// // create producer to send trades and orders data
	// producer := createProducer()

	// create the order book
	// book := engine.OrderBook{
	// 	Bids: make([]engine.Order, 0, 10000),
	// 	Asks: make([]engine.Order, 0, 10000),
	// }

	// create a channel to know when done
	// done := make(chan bool)

	// // start processing order
	// go func() {
	// 	for msg := range consumer.Messages() {
	// 		var order engine.Order
	// 		// deserialize the message
	// 		order.FromJSON(msg.Value)
	// 		// process the order
	// 		trades := book.Process(order)

	// 		log.Println("Trades length: ", len(trades))

	// 		if len(trades) != 0 {
	// 			// send trades to message queue
	// 			for _, trade := range trades {
	// 				rawTrade := trade.ToJSON()

	// 				log.Println("Publishing trade on topic -> trades")
	// 				// publish the message over receiving channel
	// 				producer.Input() <- &sarama.ProducerMessage{
	// 					Topic: "trades",
	// 					Value: sarama.ByteEncoder(rawTrade),
	// 				}
	// 			}
	// 			// mark the offset as commited
	// 			consumer.MarkOffset(msg, "")
	// 		}
	// 	}
	// 	done <- true
	// }()
	// <-done
}
