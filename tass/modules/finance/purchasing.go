package tassfinance

import tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"

// op: GetAllPurchaseOrders, path: /{cmpy_code}/finance/purchasing/purchaseorders
func (c *Client) GetAllPurchaseOrders() ([]PurchaseOrderResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddPurchaseOrder, path: /{cmpy_code}/finance/purchasing/purchaseorders
func (c *Client) AddPurchaseOrder(payload AddPurchaseOrderRequest) (PurchaseOrderResponse, error) {
	// TODO: Implementation
	return PurchaseOrderResponse{}, nil
}

// op: GetPurchaseOrderById, path: /{cmpy_code}/finance/purchasing/purchaseorders/{order_num}
func (c *Client) GetPurchaseOrderByID(orderNumber string) (PurchaseOrderResponse, error) {
	// TODO: Implementation
	return PurchaseOrderResponse{}, nil
}

// op: UpdatePurchaseOrder, path: /{cmpy_code}/finance/purchasing/purchaseorders/{order_num}
func (c *Client) UpdatePurchaseOrder(orderNumber string, payload UpdatePurchaseOrderRequest) error {
	// TODO: Implementation
	return nil
}

// op: DeletePurchaseOrder, path: /{cmpy_code}/finance/purchasing/purchaseorders/{order_num}
func (c *Client) DeletePurchaseOrder(orderNumber string) error {
	// TODO: Implementation
	return nil
}

// op: CancelPurchaseOrder, path: /{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/cancel
func (c *Client) CancelPurchaseOrder(orderNumber string) error {
	// TODO: Implementation
	return nil
}

// op: GetPurchaseOrderAttachments, path: /{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/attachments
func (c *Client) GetPurchaseOrderAttachments(orderNumber string) ([]tasscommon.FileResponse, error) {
	// TODO: Implementation
	return nil, nil
}

// op: AddPurchaseOrderAttachment, path: /{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/attachments
func (c *Client) AddPurchaseOrderAttachment(orderNumber string, payload tasscommon.FileRequest) (tasscommon.NewAttachmentResponse, error) {
	// TODO: Implementation
	return tasscommon.NewAttachmentResponse{}, nil
}

// op: DownloadPurchaseOrderAttachment, path: /{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/attachments/{attach_id}
func (c *Client) GetPurchaseOrderAttachment(orderNumber string, attachmentID string) ([]byte, error) {
	// TODO: Implementation
	return nil, nil
}

// op: DeletePurchaseOrderAttachment, path: /{cmpy_code}/finance/purchasing/purchaseorders/{order_num}/attachments/{attach_id}
func (c *Client) DeletePurchaseOrderAttachment(orderNumber string, attachmentID string) error {
	// TODO: Implementation
	return nil
}
