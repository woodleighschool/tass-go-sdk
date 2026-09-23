package tassfinance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

// CODEGEN(none): op=GetAllPurchaseOrders, path=/{cmpy_code}/finance/purchasing/purchaseorders
func (c *Client) GetAllPurchaseOrders(ctx context.Context) ([]PurchaseOrderResponse, error) {
	var result []PurchaseOrderResponse
	body, err := c.t.request(ctx, http.MethodGet, "/finance/purchasing/purchaseorders", nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddPurchaseOrder, path=/{cmpy_code}/finance/purchasing/purchaseorders
func (c *Client) AddPurchaseOrder(ctx context.Context, payload AddPurchaseOrderRequest) (PurchaseOrderResponse, error) {
	var result PurchaseOrderResponse
	if err := tasscommon.Validate(payload); err != nil {
		return PurchaseOrderResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, "/finance/purchasing/purchaseorders", nil, payload, http.StatusCreated)
	if err != nil {
		return PurchaseOrderResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return PurchaseOrderResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=GetPurchaseOrderById, path=/{cmpy_code}/finance/purchasing/purchaseorders/{order_num}
func (c *Client) GetPurchaseOrderByID(ctx context.Context, orderNumber int) (PurchaseOrderResponse, error) {
	var result PurchaseOrderResponse
	url := fmt.Sprintf("/finance/purchasing/purchaseorders/%d", orderNumber)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return PurchaseOrderResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return PurchaseOrderResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=UpdatePurchaseOrder, path=/{cmpy_code}/finance/purchasing/purchaseorders/{order_num}
func (c *Client) UpdatePurchaseOrder(ctx context.Context, orderNumber int, payload UpdatePurchaseOrderRequest) error {
	url := fmt.Sprintf("/finance/purchasing/purchaseorders/%d", orderNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return err
	}
	_, err := c.t.request(ctx, http.MethodPut, url, nil, payload, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=DeletePurchaseOrder, path=/{cmpy_code}/finance/purchasing/purchaseorders/{order_num}
func (c *Client) DeletePurchaseOrder(ctx context.Context, orderNumber int) error {
	url := fmt.Sprintf("/finance/purchasing/purchaseorders/%d", orderNumber)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=CancelPurchaseOrder, path=/{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/cancel
func (c *Client) CancelPurchaseOrder(ctx context.Context, orderNumber int) error {
	url := fmt.Sprintf("/finance/purchasing/purchaseorders/%d/cancel", orderNumber)
	_, err := c.t.request(ctx, http.MethodPost, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}

// CODEGEN(none): op=GetPurchaseOrderAttachments, path=/{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/attachments
func (c *Client) GetPurchaseOrderAttachments(ctx context.Context, orderNumber int) ([]tasscommon.FileResponse, error) {
	var result []tasscommon.FileResponse
	url := fmt.Sprintf("/finance/purchasing/purchaseorders/%d/attachments", orderNumber)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=AddPurchaseOrderAttachment, path=/{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/attachments
func (c *Client) AddPurchaseOrderAttachment(ctx context.Context, orderNumber int, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	var result tasscommon.NewAttachmentResponse
	url := fmt.Sprintf("/finance/purchasing/purchaseorders/%d/attachments", orderNumber)
	if err := tasscommon.Validate(payload); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	body, err := c.t.request(ctx, http.MethodPost, url, nil, payload, http.StatusCreated)
	if err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return tasscommon.NewAttachmentResponse{}, err
	}
	return result, nil
}

// CODEGEN(none): op=DownloadPurchaseOrderAttachment, path=/{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/attachments/{attach_id}
func (c *Client) GetPurchaseOrderAttachment(ctx context.Context, orderNumber int, attachmentID string) ([]byte, error) {
	var result []byte
	url := fmt.Sprintf("/finance/purchasing/purchaseorders/%d/attachments/%s", orderNumber, attachmentID)
	body, err := c.t.request(ctx, http.MethodGet, url, nil, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// CODEGEN(none): op=DeletePurchaseOrderAttachment, path=/{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/attachments/{attach_id}
func (c *Client) DeletePurchaseOrderAttachment(ctx context.Context, orderNumber int, attachmentID string) error {
	url := fmt.Sprintf("/finance/purchasing/purchaseorders/%d/attachments/%s", orderNumber, attachmentID)
	_, err := c.t.request(ctx, http.MethodDelete, url, nil, nil, http.StatusNoContent)
	if err != nil {
		return err
	}
	return nil
}
