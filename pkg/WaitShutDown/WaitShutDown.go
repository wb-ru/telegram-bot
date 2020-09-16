package WaitShutDown

import (
	"../models"
	"context"
	"os"
	"os/signal"
)

//Ожидание выключения, получение сигнала от пользовательских процессов
func WaitShutdown(cancel context.CancelFunc) {
	var sigs = make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	models.ChanOrderPay <- "Application Close"
	cancel()
}
