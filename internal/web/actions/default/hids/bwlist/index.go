package bwlist

import (
	bwlist_model "github.com/1uLang/zhiannet-api/hids/model/bwlist"
	"github.com/1uLang/zhiannet-api/hids/server/bwlist"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "")
}

func (this *IndexAction) RunGet(params struct {
	Address string
}) {
	defer this.Show()

	this.Data["list"] = nil
	this.Data["Address"] = ""

	//ddos节点
	list, _, err := bwlist.GetBWList(&bwlist_model.ListReq{UserId: uint64(this.UserId(true)), IP: params.Address})
	if err != nil {
		this.Data["errorMessage"] = err.Error()
		return
	}
	page := this.NewPage(int64(len(list)))
	this.Data["page"] = page.AsHTML()

	offset := page.Offset
	end := offset + page.Size
	if end > page.Total {
		end = page.Total
	}
	this.Data["list"] = list[offset:end]
	this.Data["Address"] = params.Address

}
