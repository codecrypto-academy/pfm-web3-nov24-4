package handlers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/codecrypto-academy/pfm-web3-nov24-4/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetFabricConnectionChain(c *gin.Context) {
	var params models.ConnectionChain

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "inspect", "--output", "configs/"+params.MSPID+"/"+params.Output, "-o", params.MSPID)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	log.Debugf("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))

	c.JSON(http.StatusOK, gin.H{
		"message": "Fabric Connection Chain obtenido exitosamente",
		"output":  string(output),
	})
}

func GetFabricUserCerts(c *gin.Context) {
	var params models.GetUserCertsOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "ca", "enroll",
		"--name="+params.CAName,
		"--user="+params.User,
		"--secret="+params.Secret,
		"--mspid="+params.MSPID,
		"--output="+"configs/"+params.MSPID+"/"+params.Output)

	log.Debugf("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Certificado del usuario %s guardado en configs/%s.yaml", params.User, params.Output),
	})
}

func AttachUserToConnection(c *gin.Context) {
	var params models.AttachUserOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "utils", "adduser",
		"--userPath="+"configs/"+params.MSPID+"/"+params.UserPath,
		"--config="+"configs/"+params.MSPID+"/"+params.Config,
		"--username="+params.Username,
		"--mspid="+params.MSPID)

	log.Debugf("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Usuario %s añadido a la cadena de conexión correctamente", params.Username),
	})
}

func DeleteOrdererOrganization(c *gin.Context) {
	var params models.DeleteOrdererOrganizationOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "ca", "delete",
		"--name="+params.CAName,
		"--namespace="+params.Namespace)

	log.Debugf("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	for _, orderer := range params.Orderers {
		cmd = exec.Command("kubectl", "hlf", "ordnode", "delete",
			"--name="+orderer,
			"--namespace="+params.Namespace)

		output, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
			return
		}

		log.Debugf("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Organización eliminada correctamente",
	})
}

func DeletePeerOrganization(c *gin.Context) {
	var params models.DeletePeerOrganizationOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "ca", "delete",
		"--name="+params.CAName,
		"--namespace="+params.Namespace)

	log.Debugf("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		// return
	}

	for _, peer := range params.Peers {
		cmd = exec.Command("kubectl", "hlf", "peer", "delete",
			"--name="+peer,
			"--namespace="+params.Namespace)

		output, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
			return
		}

		log.Debugf("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Organización eliminada correctamente",
	})
}

func CreateFabricWallet(c *gin.Context) {
	var params models.CreateWalletOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "create", "secret", "generic",
		params.WalletName,
		"--namespace="+params.Namespace,
	)

	for _, fromFile := range params.FromFiles {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--from-file=configs/%s", fromFile))
	}

	log.Debugf("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet created successfully",
	})
}
