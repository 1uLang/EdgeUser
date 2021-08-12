package databackup

import (
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type CreateAction struct {
	actionutils.ParentAction
}

func (this *CreateAction) Init() {
	this.Nav("", "", "create")
}

func (this *CreateAction) RunGet(params struct{}) {

	this.Show()
}

func (this *CreateAction) RunPost(params struct {
	DomainNames     []string
	Protocols       []string
	CertIdsJSON     []byte
	OriginsJSON     []byte
	RequestHostType int32
	RequestHost     string
	CacheCondsJSON  []byte

	//Must *actions.Must
	//CSRF *actionutils.CSRF
}) {

	this.Success()
}
