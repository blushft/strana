package entity

import (
	"context"

	"github.com/blushft/strana/modules/sink/reporter/store"
	"github.com/blushft/strana/modules/sink/reporter/store/ent"
)

type Device struct {
	ID           string                 `json:"id"`
	Manufacturer string                 `json:"manufacturer"`
	Model        string                 `json:"model"`
	Name         string                 `json:"name"`
	Type         string                 `json:"type"`
	Version      string                 `json:"version"`
	Mobile       bool                   `json:"mobile"`
	Tablet       bool                   `json:"tablet"`
	Desktop      bool                   `json:"desktop"`
	Properties   map[string]interface{} `json:"properties"`
}

type DeviceReader interface {
	List(QueryParams) ([]*Device, error)
	Get(id string) (*Device, error)
}

type DeviceWriter interface {
	Create(*Device) error
	Update(*Device) error
	Delete(*Device) error
}

type deviceRepo interface {
	DeviceReader
	DeviceWriter
}

type DeviceReporter interface {
	DeviceReader
}

type DeviceManager interface {
	deviceRepo
}

type deviceManager struct {
	store *store.Store
}

func NewDeviceService(s *store.Store) *deviceManager {
	return &deviceManager{
		store: s,
	}
}

func (mgr *deviceManager) List(qp QueryParams) ([]*Device, error) {
	c := mgr.store.Client().Device
	q := c.Query()

	switch {
	case qp.Limit == 0:
		q.Limit(DefaultLimit)
	case qp.Limit > 0:
		q.Limit(qp.Limit)
	}

	if qp.Offset > 0 && qp.Limit > -1 {
		q.Offset(qp.Offset)
	}

	recs, err := q.All(context.TODO())
	if err != nil {
		return nil, err
	}

	res := make([]*Device, 0, len(recs))
	for _, rec := range recs {
		res = append(res, deviceSchemaToEntity(rec))
	}

	return res, nil
}

func (mgr *deviceManager) Get(id string) (*Device, error) {
	c := mgr.store.Client().Device

	rec, err := c.Get(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	return deviceSchemaToEntity(rec), nil
}

func (mgr *deviceManager) Create(e *Device) error {
	c := mgr.store.Client().Device

	_, err := deviceEntityCreate(c, e).Save(context.TODO())

	return err
}

func (mgr *deviceManager) Update(e *Device) error {
	c := mgr.store.Client().Device

	_, err := deviceEntityUpdate(c, e).Save(context.TODO())

	return err
}

func (mgr *deviceManager) Delete(e *Device) error {
	c := mgr.store.Client().Device

	return c.DeleteOneID(e.ID).Exec(context.TODO())
}

func deviceEntityCreate(c *ent.DeviceClient, e *Device) *ent.DeviceCreate {
	return c.Create().
		SetID(e.ID).
		SetManufacturer(e.Manufacturer).
		SetModel(e.Model).
		SetName(e.Name).
		SetType(e.Type).
		SetVersion(e.Version).
		SetMobile(e.Mobile).
		SetTablet(e.Tablet).
		SetDesktop(e.Desktop).
		SetProperties(e.Properties)
}

func deviceEntityUpdate(c *ent.DeviceClient, e *Device) *ent.DeviceUpdate {
	return c.Update().
		SetManufacturer(e.Manufacturer).
		SetModel(e.Model).
		SetName(e.Name).
		SetType(e.Type).
		SetVersion(e.Version).
		SetMobile(e.Mobile).
		SetTablet(e.Tablet).
		SetDesktop(e.Desktop).
		SetProperties(e.Properties)
}

func deviceSchemaToEntity(sch *ent.Device) *Device {
	return &Device{
		ID:           sch.ID,
		Manufacturer: sch.Manufacturer,
		Model:        sch.Model,
		Name:         sch.Name,
		Type:         sch.Type,
		Version:      sch.Version,
		Mobile:       sch.Mobile,
		Tablet:       sch.Tablet,
		Desktop:      sch.Desktop,
		Properties:   sch.Properties,
	}
}
