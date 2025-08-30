package models

import (
	"errors"
	"slices"

	"github.com/google/uuid"
)

type Pixel struct {
	BaseModel
	Name          string     `json:"name" gorm:"type:varchar(255);not null"`
	Description   string     `json:"description" gorm:"type:text"`
	Audience      int        `json:"audience" gorm:"type:integer;default:0"`
	WebsiteLogo   string     `json:"website_logo" gorm:"type:varchar(255)"`
	WebsiteUrl    string     `json:"website_url" gorm:"type:varchar(255)"`
	Price         float64    `json:"price" gorm:"type:float(10,2)"`
	Currency      string     `json:"currency" gorm:"type:varchar(10)"`
	AudienceProof string     `json:"audience_proof" gorm:"type:varchar(255)"`
	CreatorId     *uuid.UUID `json:"creator_id"`
	Creator       *User      `json:"creator" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReviewerId    *uuid.UUID `json:"reviewer_id"`
	Reviewer      *User      `json:"reviewer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SourceId      *uuid.UUID `json:"source_id"`
	Source        *Source    `json:"source" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EventId       *uuid.UUID `json:"event_id"`
	Event         *Event     `json:"event" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TypeId        *uuid.UUID `json:"type_id"`
	Type          *Type      `json:"type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StatusId      *uuid.UUID `json:"status_id"`
	Status        *Status    `json:"status" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryID    *uuid.UUID `json:"category_id"`
	Category      *Category  `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

var (
	// Review Process Statuses
	PixelStatusUnderReview      = uuid.MustParse("22d6fd06-4cdb-4ed7-bab7-7c608c2f182e")
	PixelStatusApproved         = uuid.MustParse("43df7b00-620c-4983-9da9-4da73da0fcf3")
	PixelStatusRejected         = uuid.MustParse("1033a937-a743-431b-88cb-5a2c4ffade96")
	PixelStatusDraft            = uuid.MustParse("71b14051-8b76-4224-b122-5073e4383bb9")
	PixelStatusRevisionRequired = uuid.MustParse("4167084e-4a4f-4b36-9418-2073999379ce")

	// Availability Statuses
	PixelStatusAvailable  = uuid.MustParse("8c91705a-11d0-4461-816b-d3775a60d967")
	PixelStatusOutOfStock = uuid.MustParse("eac9b264-ac0a-4b4c-aa43-1a9ff9ba422a")
	PixelStatusSold       = uuid.MustParse("85cc0b34-f748-4311-80a9-e891e7cec14b")
	PixelStatusReserved   = uuid.MustParse("a60adbb4-88ac-4b64-9757-e0d999f664be")
	PixelStatusSuspended  = uuid.MustParse("740db8a7-9f0c-4f79-9581-7f7025990c1c")
)

func (p *Pixel) IsApproved() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusApproved
}

func (p *Pixel) IsAvailable() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusAvailable
}

func (p *Pixel) IsUnderReview() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusUnderReview
}

func (p *Pixel) IsRejected() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusRejected
}

func (p *Pixel) IsSold() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusSold
}

func (p *Pixel) IsDraft() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusDraft
}

func (p *Pixel) IsOutOfStock() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusOutOfStock
}

func (p *Pixel) IsReserved() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusReserved
}

func (p *Pixel) IsSuspended() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusSuspended
}

func (p *Pixel) NeedsRevision() bool {
	return p.StatusId != nil && *p.StatusId == PixelStatusRevisionRequired
}

func (p *Pixel) SetUnderReview() {
	p.StatusId = &PixelStatusUnderReview
}

func (p *Pixel) SetApproved() {
	p.StatusId = &PixelStatusApproved
}

func (p *Pixel) SetRejected() {
	p.StatusId = &PixelStatusRejected
}

func (p *Pixel) SetAvailable() {
	p.StatusId = &PixelStatusAvailable
}

func (p *Pixel) SetSold() {
	p.StatusId = &PixelStatusSold
}

func (p *Pixel) SetReserved() {
	p.StatusId = &PixelStatusReserved
}

func (p *Pixel) SetOutOfStock() {
	p.StatusId = &PixelStatusOutOfStock
}

func (p *Pixel) SetSuspended() {
	p.StatusId = &PixelStatusSuspended
}

func (p *Pixel) SetDraft() {
	p.StatusId = &PixelStatusDraft
}

func (p *Pixel) SetRevisionRequired() {
	p.StatusId = &PixelStatusRevisionRequired
}

func (p *Pixel) CanBeApproved() bool {
	return p.IsUnderReview() || p.NeedsRevision()
}

func (p *Pixel) CanBePublished() bool {
	return p.IsApproved()
}

func (p *Pixel) CanBePurchased() bool {
	return p.IsAvailable()
}

func (p *Pixel) CanBeEdited() bool {
	return p.IsDraft() || p.IsRejected() || p.NeedsRevision()
}

func (p *Pixel) CanTransitionTo(newStatusId uuid.UUID) bool {
	if p.StatusId == nil {
		return newStatusId == PixelStatusDraft
	}

	currentStatus := *p.StatusId

	validTransitions := map[uuid.UUID][]uuid.UUID{
		PixelStatusDraft: {
			PixelStatusUnderReview,
		},
		PixelStatusUnderReview: {
			PixelStatusApproved,
			PixelStatusRejected,
			PixelStatusRevisionRequired,
		},
		PixelStatusRevisionRequired: {
			PixelStatusUnderReview,
			PixelStatusDraft,
		},
		PixelStatusRejected: {
			PixelStatusDraft,
			PixelStatusUnderReview,
		},
		PixelStatusApproved: {
			PixelStatusAvailable,
			PixelStatusSuspended,
		},
		PixelStatusAvailable: {
			PixelStatusReserved,
			PixelStatusSold,
			PixelStatusOutOfStock,
			PixelStatusSuspended,
		},
		PixelStatusReserved: {
			PixelStatusSold,
			PixelStatusAvailable,
		},
		PixelStatusOutOfStock: {
			PixelStatusAvailable,
		},
		PixelStatusSuspended: {
			PixelStatusDraft,
			PixelStatusAvailable,
		},
	}

	allowedTransitions, exists := validTransitions[currentStatus]
	if !exists {
		return false
	}

	return slices.Contains(allowedTransitions, newStatusId)
}

func (p *Pixel) TransitionTo(newStatusId uuid.UUID) error {
	if !p.CanTransitionTo(newStatusId) {
		return errors.New("invalid status transition")
	}
	p.StatusId = &newStatusId
	return nil
}

func (p *Pixel) GetStatusInfo() (name, description, color string) {
	if p.StatusId == nil {
		return "Unknown", "Status not set", "#6B7280"
	}

	statusMap := map[uuid.UUID]struct {
		name, description, color string
	}{
		PixelStatusUnderReview:      {"Under Review", "Pixel is currently being reviewed by administrators", "#F59E0B"},
		PixelStatusApproved:         {"Approved", "Pixel has been approved and is ready for listing", "#10B981"},
		PixelStatusRejected:         {"Rejected", "Pixel has been rejected and requires modifications", "#EF4444"},
		PixelStatusDraft:            {"Draft", "Pixel is in draft state and not yet submitted for review", "#6B7280"},
		PixelStatusRevisionRequired: {"Revision Required", "Pixel needs revisions before approval", "#F97316"},
		PixelStatusAvailable:        {"Available", "Pixel is available for purchase", "#059669"},
		PixelStatusOutOfStock:       {"Out of Stock", "Pixel is temporarily unavailable", "#DC2626"},
		PixelStatusSold:             {"Sold", "Pixel has been sold and is no longer available", "#7C2D12"},
		PixelStatusReserved:         {"Reserved", "Pixel is reserved for a specific buyer", "#7C3AED"},
		PixelStatusSuspended:        {"Suspended", "Pixel listing has been suspended", "#991B1B"},
	}

	if info, exists := statusMap[*p.StatusId]; exists {
		return info.name, info.description, info.color
	}

	return "Unknown", "Unknown status", "#6B7280"
}
