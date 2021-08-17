package bills

import (
	"fmt"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeUser/internal/web/actions/actionutils"
	"github.com/iwind/TeaGo/maps"
)

type IndexAction struct {
	actionutils.ParentAction
}

func (this *IndexAction) Init() {
	this.Nav("", "", "index")
}

func (this *IndexAction) RunGet(params struct {
	PaidFlag int32 `default:"-1"`
	Month    string
}) {
	if !this.ValidateFeature("finance") {
		return
	}

	countResp, err := this.RPC().UserBillRPC().CountAllUserBills(this.UserContext(), &pb.CountAllUserBillsRequest{})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	page := this.NewPage(countResp.Count)
	this.Data["page"] = page.AsHTML()

	billsResp, err := this.RPC().UserBillRPC().ListUserBills(this.UserContext(), &pb.ListUserBillsRequest{
		PaidFlag: params.PaidFlag,
		UserId:   this.UserId(true),
		Month:    params.Month,
		Offset:   page.Offset,
		Size:     page.Size,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}
	billMaps := []maps.Map{}
	for _, bill := range billsResp.UserBills {
		var userMap maps.Map = nil
		if bill.User != nil {
			userMap = maps.Map{
				"id":       bill.User.Id,
				"fullname": bill.User.Fullname,
			}
		}
		billMaps = append(billMaps, maps.Map{
			"id":          bill.Id,
			"isPaid":      bill.IsPaid,
			"month":       bill.Month,
			"amount":      fmt.Sprintf("%.2f", bill.Amount),
			"typeName":    bill.TypeName,
			"user":        userMap,
			"description": bill.Description,
		})
	}
	this.Data["bills"] = billMaps

	this.Show()
}
