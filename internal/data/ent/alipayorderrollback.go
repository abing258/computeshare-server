// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/alipayorderrollback"
)

// AlipayOrderRollback is the model entity for the AlipayOrderRollback schema.
type AlipayOrderRollback struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 通知校验ID
	NotifyID string `json:"notify_id,omitempty"`
	// 通知类型
	NotifyType string `json:"notify_type,omitempty"`
	// 通知时间
	NotifyTime string `json:"notify_time,omitempty"`
	// 编码格式，如 utf-8、gbk、gb2312 等
	Charset string `json:"charset,omitempty"`
	// 调用的接口版本，固定为：1.0
	Version string `json:"version,omitempty"`
	// 签名类型
	SignType string `json:"sign_type,omitempty"`
	// 签名
	Sign string `json:"sign,omitempty"`
	// 支付成功的各个渠道金额信息。详情可查看 资金明细信息说明
	FundBillList string `json:"fund_bill_list,omitempty"`
	// 实收金额
	ReceiptAmount string `json:"receipt_amount,omitempty"`
	// 用户在交易中支付的可开发票的金额
	InvoiceAmount string `json:"invoice_amount,omitempty"`
	// 付款金额
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty"`
	// 集分宝金额
	PointAmount string `json:"point_amount,omitempty"`
	// 本交易支付时所有优惠券信息，详情可查看 优惠券信息说明
	VoucherDetailList string `json:"voucher_detail_list,omitempty"`
	// 公共回传参数，如果请求时传递了该参数，则返回给商家时会在异步通知时将该参数原样返回。本参数必须进行 UrlEncode 之后才可以发送给支付宝。
	PassbackParams string `json:"passback_params,omitempty"`
	// 支付宝交易号
	TradeNo string `json:"trade_no,omitempty"`
	// 开发者id
	AppID string `json:"app_id,omitempty"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 商户业务号
	OutBizNo string `json:"out_biz_no,omitempty"`
	// 买家支付宝ID
	BuyerID string `json:"buyer_id,omitempty"`
	// 卖家支付宝id
	SellerID string `json:"seller_id,omitempty"`
	// 交易状态
	TradeStatus string `json:"trade_status,omitempty"`
	// 订单金额
	TotalAmount string `json:"total_amount,omitempty"`
	// 总退款金额
	RefundFee string `json:"refund_fee,omitempty"`
	// 订单标题
	Subject string `json:"subject,omitempty"`
	// 订单的备注、描述、明细等。对应请求时的 body 参数，原样通知回来
	Body string `json:"body,omitempty"`
	// 交易创建时间
	GmtCreate string `json:"gmt_create,omitempty"`
	// 交易付款时间
	GmtPayment string `json:"gmt_payment,omitempty"`
	// 交易关闭时间
	GmtClose string `json:"gmt_close,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AlipayOrderRollback) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case alipayorderrollback.FieldID:
			values[i] = new(sql.NullInt64)
		case alipayorderrollback.FieldNotifyID, alipayorderrollback.FieldNotifyType, alipayorderrollback.FieldNotifyTime, alipayorderrollback.FieldCharset, alipayorderrollback.FieldVersion, alipayorderrollback.FieldSignType, alipayorderrollback.FieldSign, alipayorderrollback.FieldFundBillList, alipayorderrollback.FieldReceiptAmount, alipayorderrollback.FieldInvoiceAmount, alipayorderrollback.FieldBuyerPayAmount, alipayorderrollback.FieldPointAmount, alipayorderrollback.FieldVoucherDetailList, alipayorderrollback.FieldPassbackParams, alipayorderrollback.FieldTradeNo, alipayorderrollback.FieldAppID, alipayorderrollback.FieldOutTradeNo, alipayorderrollback.FieldOutBizNo, alipayorderrollback.FieldBuyerID, alipayorderrollback.FieldSellerID, alipayorderrollback.FieldTradeStatus, alipayorderrollback.FieldTotalAmount, alipayorderrollback.FieldRefundFee, alipayorderrollback.FieldSubject, alipayorderrollback.FieldBody, alipayorderrollback.FieldGmtCreate, alipayorderrollback.FieldGmtPayment, alipayorderrollback.FieldGmtClose:
			values[i] = new(sql.NullString)
		case alipayorderrollback.FieldCreateTime, alipayorderrollback.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AlipayOrderRollback fields.
func (aor *AlipayOrderRollback) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case alipayorderrollback.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aor.ID = int(value.Int64)
		case alipayorderrollback.FieldNotifyID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notify_id", values[i])
			} else if value.Valid {
				aor.NotifyID = value.String
			}
		case alipayorderrollback.FieldNotifyType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notify_type", values[i])
			} else if value.Valid {
				aor.NotifyType = value.String
			}
		case alipayorderrollback.FieldNotifyTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notify_time", values[i])
			} else if value.Valid {
				aor.NotifyTime = value.String
			}
		case alipayorderrollback.FieldCharset:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field charset", values[i])
			} else if value.Valid {
				aor.Charset = value.String
			}
		case alipayorderrollback.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				aor.Version = value.String
			}
		case alipayorderrollback.FieldSignType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sign_type", values[i])
			} else if value.Valid {
				aor.SignType = value.String
			}
		case alipayorderrollback.FieldSign:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sign", values[i])
			} else if value.Valid {
				aor.Sign = value.String
			}
		case alipayorderrollback.FieldFundBillList:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fund_bill_list", values[i])
			} else if value.Valid {
				aor.FundBillList = value.String
			}
		case alipayorderrollback.FieldReceiptAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receipt_amount", values[i])
			} else if value.Valid {
				aor.ReceiptAmount = value.String
			}
		case alipayorderrollback.FieldInvoiceAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invoice_amount", values[i])
			} else if value.Valid {
				aor.InvoiceAmount = value.String
			}
		case alipayorderrollback.FieldBuyerPayAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field buyer_pay_amount", values[i])
			} else if value.Valid {
				aor.BuyerPayAmount = value.String
			}
		case alipayorderrollback.FieldPointAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field point_amount", values[i])
			} else if value.Valid {
				aor.PointAmount = value.String
			}
		case alipayorderrollback.FieldVoucherDetailList:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field voucher_detail_list", values[i])
			} else if value.Valid {
				aor.VoucherDetailList = value.String
			}
		case alipayorderrollback.FieldPassbackParams:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field passback_params", values[i])
			} else if value.Valid {
				aor.PassbackParams = value.String
			}
		case alipayorderrollback.FieldTradeNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field trade_no", values[i])
			} else if value.Valid {
				aor.TradeNo = value.String
			}
		case alipayorderrollback.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				aor.AppID = value.String
			}
		case alipayorderrollback.FieldOutTradeNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field out_trade_no", values[i])
			} else if value.Valid {
				aor.OutTradeNo = value.String
			}
		case alipayorderrollback.FieldOutBizNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field out_biz_no", values[i])
			} else if value.Valid {
				aor.OutBizNo = value.String
			}
		case alipayorderrollback.FieldBuyerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field buyer_id", values[i])
			} else if value.Valid {
				aor.BuyerID = value.String
			}
		case alipayorderrollback.FieldSellerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seller_id", values[i])
			} else if value.Valid {
				aor.SellerID = value.String
			}
		case alipayorderrollback.FieldTradeStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field trade_status", values[i])
			} else if value.Valid {
				aor.TradeStatus = value.String
			}
		case alipayorderrollback.FieldTotalAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value.Valid {
				aor.TotalAmount = value.String
			}
		case alipayorderrollback.FieldRefundFee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refund_fee", values[i])
			} else if value.Valid {
				aor.RefundFee = value.String
			}
		case alipayorderrollback.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				aor.Subject = value.String
			}
		case alipayorderrollback.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				aor.Body = value.String
			}
		case alipayorderrollback.FieldGmtCreate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gmt_create", values[i])
			} else if value.Valid {
				aor.GmtCreate = value.String
			}
		case alipayorderrollback.FieldGmtPayment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gmt_payment", values[i])
			} else if value.Valid {
				aor.GmtPayment = value.String
			}
		case alipayorderrollback.FieldGmtClose:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gmt_close", values[i])
			} else if value.Valid {
				aor.GmtClose = value.String
			}
		case alipayorderrollback.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				aor.CreateTime = value.Time
			}
		case alipayorderrollback.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				aor.UpdateTime = value.Time
			}
		default:
			aor.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AlipayOrderRollback.
// This includes values selected through modifiers, order, etc.
func (aor *AlipayOrderRollback) Value(name string) (ent.Value, error) {
	return aor.selectValues.Get(name)
}

// Update returns a builder for updating this AlipayOrderRollback.
// Note that you need to call AlipayOrderRollback.Unwrap() before calling this method if this AlipayOrderRollback
// was returned from a transaction, and the transaction was committed or rolled back.
func (aor *AlipayOrderRollback) Update() *AlipayOrderRollbackUpdateOne {
	return NewAlipayOrderRollbackClient(aor.config).UpdateOne(aor)
}

// Unwrap unwraps the AlipayOrderRollback entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aor *AlipayOrderRollback) Unwrap() *AlipayOrderRollback {
	_tx, ok := aor.config.driver.(*txDriver)
	if !ok {
		panic("ent: AlipayOrderRollback is not a transactional entity")
	}
	aor.config.driver = _tx.drv
	return aor
}

// String implements the fmt.Stringer.
func (aor *AlipayOrderRollback) String() string {
	var builder strings.Builder
	builder.WriteString("AlipayOrderRollback(")
	builder.WriteString(fmt.Sprintf("id=%v, ", aor.ID))
	builder.WriteString("notify_id=")
	builder.WriteString(aor.NotifyID)
	builder.WriteString(", ")
	builder.WriteString("notify_type=")
	builder.WriteString(aor.NotifyType)
	builder.WriteString(", ")
	builder.WriteString("notify_time=")
	builder.WriteString(aor.NotifyTime)
	builder.WriteString(", ")
	builder.WriteString("charset=")
	builder.WriteString(aor.Charset)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(aor.Version)
	builder.WriteString(", ")
	builder.WriteString("sign_type=")
	builder.WriteString(aor.SignType)
	builder.WriteString(", ")
	builder.WriteString("sign=")
	builder.WriteString(aor.Sign)
	builder.WriteString(", ")
	builder.WriteString("fund_bill_list=")
	builder.WriteString(aor.FundBillList)
	builder.WriteString(", ")
	builder.WriteString("receipt_amount=")
	builder.WriteString(aor.ReceiptAmount)
	builder.WriteString(", ")
	builder.WriteString("invoice_amount=")
	builder.WriteString(aor.InvoiceAmount)
	builder.WriteString(", ")
	builder.WriteString("buyer_pay_amount=")
	builder.WriteString(aor.BuyerPayAmount)
	builder.WriteString(", ")
	builder.WriteString("point_amount=")
	builder.WriteString(aor.PointAmount)
	builder.WriteString(", ")
	builder.WriteString("voucher_detail_list=")
	builder.WriteString(aor.VoucherDetailList)
	builder.WriteString(", ")
	builder.WriteString("passback_params=")
	builder.WriteString(aor.PassbackParams)
	builder.WriteString(", ")
	builder.WriteString("trade_no=")
	builder.WriteString(aor.TradeNo)
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(aor.AppID)
	builder.WriteString(", ")
	builder.WriteString("out_trade_no=")
	builder.WriteString(aor.OutTradeNo)
	builder.WriteString(", ")
	builder.WriteString("out_biz_no=")
	builder.WriteString(aor.OutBizNo)
	builder.WriteString(", ")
	builder.WriteString("buyer_id=")
	builder.WriteString(aor.BuyerID)
	builder.WriteString(", ")
	builder.WriteString("seller_id=")
	builder.WriteString(aor.SellerID)
	builder.WriteString(", ")
	builder.WriteString("trade_status=")
	builder.WriteString(aor.TradeStatus)
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(aor.TotalAmount)
	builder.WriteString(", ")
	builder.WriteString("refund_fee=")
	builder.WriteString(aor.RefundFee)
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(aor.Subject)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(aor.Body)
	builder.WriteString(", ")
	builder.WriteString("gmt_create=")
	builder.WriteString(aor.GmtCreate)
	builder.WriteString(", ")
	builder.WriteString("gmt_payment=")
	builder.WriteString(aor.GmtPayment)
	builder.WriteString(", ")
	builder.WriteString("gmt_close=")
	builder.WriteString(aor.GmtClose)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(aor.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(aor.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AlipayOrderRollbacks is a parsable slice of AlipayOrderRollback.
type AlipayOrderRollbacks []*AlipayOrderRollback
