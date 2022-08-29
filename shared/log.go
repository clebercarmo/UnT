package shared

import (
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


// SetLog Grava dado de log estruturado no arquivo info.log
// Tipos -
// Info: Informações relevantes - 
// Error: Impedimentos, Excessões
func SetLog(tipo, mensagem string){


	writeSyncer, _ := os.OpenFile("./info.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)					
	encoderConfig := zap.NewProductionEncoderConfig()					 
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)					
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)	 
	log := zap.New(core,zap.AddCaller())		
	
	
	if tipo == "Info"{
		log.Info(mensagem)

	}

	if tipo == "Error"{
		log.Error(mensagem)
	}

	
	

}