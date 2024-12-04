package models

type ConnectionChain struct {
	Output string `json:"output"`
	MSPID  string `json:"mspid"`
}

type AddUserToConnectionChainOptions struct {
	UserPath string `json:"userPath"`
	Config   string `json:"config"`
	Username string `json:"username"`
	MSPID    string `json:"mspid"`
}

type GetUserCertsOptions struct {
	Name   string `json:"name"`
	User   string `json:"user"`
	Secret string `json:"secret"`
	MSPID  string `json:"mspid"`
	CAName string `json:"caName"`
	Output string `json:"output"`
}

type AttachUserOptions struct {
	UserPath string `json:"userpath"`
	Config   string `json:"config"`
	Username string `json:"username"`
	MSPID    string `json:"mspid"`
}

type DeleteOrdererOrganizationOptions struct {
	CAName    string   `json:"caName"`
	Orderers  []string `json:"orderers"`
	Namespace string   `json:"namespace"`
}

type DeletePeerOrganizationOptions struct {
	CAName    string   `json:"ca_name"`
	Peers     []string `json:"peers"`
	Namespace string   `json:"namespace"`
}

type CreateWalletOptions struct {
	WalletName string   `json:"wallet_name"`
	Namespace  string   `json:"namespace"`
	FromFiles  []string `json:"from_files"`
}
