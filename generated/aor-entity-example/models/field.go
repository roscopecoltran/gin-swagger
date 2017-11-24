// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Field field
// swagger:model Field
type Field struct {

	// choices
	Choices []IChoiceItem `json:"choices"`

	// component
	Component string `json:"component,omitempty"`

	// data source Id
	DataSourceID string `json:"dataSourceId,omitempty" gorm:"data_source_id" yaml:"data_source_id" toml:"data_source_id"`

	// default value
	DefaultValue string `json:"defaultValue,omitempty" gorm:"default_value" yaml:"default_value" toml:"default_value"`

	// eid
	Eid string `json:"eid,omitempty" gorm:"eid" yaml:"eid" toml:"eid"`

	// id
	ID string `json:"id,omitempty" gorm:"primary_key" yaml:"-" toml:"-"`

	// input type
	InputType string `json:"inputType,omitempty" gorm:"input_type" yaml:"input_type" toml:"input_type"`

	// is auto incremented
	IsAutoIncremented bool `json:"isAutoIncremented,omitempty" gorm:"is_auto_incremented" yaml:"is_auto_incremented" toml:"is_auto_incremented"`

	// is part of primary key
	IsPartOfPrimaryKey bool `json:"isPartOfPrimaryKey,omitempty" gorm:"is_part_of_primary_key" yaml:"is_part_of_primary_key" toml:"is_part_of_primary_key"`

	// label
	Label string `json:"label,omitempty" gorm:"label" yaml:"label" toml:"label"`

	// max length
	MaxLength int32 `json:"maxLength,omitempty" gorm:"max_length" yaml:"max_length" toml:"max_length"`

	// max value
	MaxValue string `json:"maxValue,omitempty" gorm:"max_value" yaml:"max_value" toml:"max_value"`

	// min value
	MinValue string `json:"minValue,omitempty" gorm:"min_value" yaml:"min_value" toml:"min_value"`

	// name
	Name string `json:"name,omitempty" gorm:"name" yaml:"name" toml:"name"`

	// reference
	Reference string `json:"reference,omitempty" gorm:"reference" yaml:"reference" toml:"reference"`

	// reference option text
	ReferenceOptionText string `json:"referenceOptionText,omitempty" gorm:"reference_option_text" yaml:"reference_option_text" toml:"reference_option_text"`

	// required
	Required bool `json:"required,omitempty" gorm:"required" yaml:"required" toml:"required"`

	// show in create
	ShowInCreate bool `json:"showInCreate,omitempty" gorm:"show_in_create" yaml:"show_in_create" toml:"show_in_create"`

	// show in edit
	ShowInEdit bool `json:"showInEdit,omitempty" gorm:"show_in_edit" yaml:"show_in_edit" toml:"show_in_edit"`

	// show in filter
	ShowInFilter bool `json:"showInFilter,omitempty" gorm:"show_in_filter" yaml:"show_in_filter" toml:"show_in_filter"`

	// show in list
	ShowInList bool `json:"showInList,omitempty" gorm:"show_in_list" yaml:"show_in_list" toml:"show_in_list"`

	// show in show
	ShowInShow bool `json:"showInShow,omitempty" gorm:"show_in_show" yaml:"show_in_show" toml:"show_in_show"`
}

// Validate validates this field
func (m *Field) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChoices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInputType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Field) validateChoices(formats strfmt.Registry) error {

	if swag.IsZero(m.Choices) { // not required
		return nil
	}

	for i := 0; i < len(m.Choices); i++ {

	}

	return nil
}

var fieldTypeComponentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Autocomplete","Boolean","NullableBoolean","CheckboxGroup","Date","Disabled","File","Image","LongText","Number","RadioButtonGroup","ReferenceArray","Reference","RichText","SelectArray","Select","Text"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldTypeComponentPropEnum = append(fieldTypeComponentPropEnum, v)
	}
}

const (
	// FieldComponentAutocomplete captures enum value "Autocomplete"
	FieldComponentAutocomplete string = "Autocomplete"
	// FieldComponentBoolean captures enum value "Boolean"
	FieldComponentBoolean string = "Boolean"
	// FieldComponentNullableBoolean captures enum value "NullableBoolean"
	FieldComponentNullableBoolean string = "NullableBoolean"
	// FieldComponentCheckboxGroup captures enum value "CheckboxGroup"
	FieldComponentCheckboxGroup string = "CheckboxGroup"
	// FieldComponentDate captures enum value "Date"
	FieldComponentDate string = "Date"
	// FieldComponentDisabled captures enum value "Disabled"
	FieldComponentDisabled string = "Disabled"
	// FieldComponentFile captures enum value "File"
	FieldComponentFile string = "File"
	// FieldComponentImage captures enum value "Image"
	FieldComponentImage string = "Image"
	// FieldComponentLongText captures enum value "LongText"
	FieldComponentLongText string = "LongText"
	// FieldComponentNumber captures enum value "Number"
	FieldComponentNumber string = "Number"
	// FieldComponentRadioButtonGroup captures enum value "RadioButtonGroup"
	FieldComponentRadioButtonGroup string = "RadioButtonGroup"
	// FieldComponentReferenceArray captures enum value "ReferenceArray"
	FieldComponentReferenceArray string = "ReferenceArray"
	// FieldComponentReference captures enum value "Reference"
	FieldComponentReference string = "Reference"
	// FieldComponentRichText captures enum value "RichText"
	FieldComponentRichText string = "RichText"
	// FieldComponentSelectArray captures enum value "SelectArray"
	FieldComponentSelectArray string = "SelectArray"
	// FieldComponentSelect captures enum value "Select"
	FieldComponentSelect string = "Select"
	// FieldComponentText captures enum value "Text"
	FieldComponentText string = "Text"
)

// prop value enum
func (m *Field) validateComponentEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fieldTypeComponentPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Field) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	// value enum
	if err := m.validateComponentEnum("component", "body", m.Component); err != nil {
		return err
	}

	return nil
}

var fieldTypeInputTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text","email","password","url"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldTypeInputTypePropEnum = append(fieldTypeInputTypePropEnum, v)
	}
}

const (
	// FieldInputTypeText captures enum value "text"
	FieldInputTypeText string = "text"
	// FieldInputTypeEmail captures enum value "email"
	FieldInputTypeEmail string = "email"
	// FieldInputTypePassword captures enum value "password"
	FieldInputTypePassword string = "password"
	// FieldInputTypeURL captures enum value "url"
	FieldInputTypeURL string = "url"
)

// prop value enum
func (m *Field) validateInputTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, fieldTypeInputTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Field) validateInputType(formats strfmt.Registry) error {

	if swag.IsZero(m.InputType) { // not required
		return nil
	}

	// value enum
	if err := m.validateInputTypeEnum("inputType", "body", m.InputType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Field) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Field) UnmarshalBinary(b []byte) error {
	var res Field
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
