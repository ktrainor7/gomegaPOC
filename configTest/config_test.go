// config_test.go

package config

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//. "."
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config", func() {
	lineItemA := LineItem{ID: "lineItemA", Name: "Line Item A", Price: 10.20, Qty: 0}
	lineItemB := LineItem{ID: "lineItemB", Name: "Line Item B", Price: 7.66, Qty: 0}

	Context("initially", func() {
		config := Config{}

		It("has 0 line items", func() {
			Expect(config.TotalUniqueLineItems()).Should(BeZero())
		})

		It("has 0 units", func() {
			Expect(config.TotalUnits()).Should(BeZero())
		})

		Specify("the total amount is 0.00", func() {
			Expect(config.TotalAmount()).Should(BeZero())
		})
	})

	Context("when a new line item is added", func() {
		config := Config{}

		originalLineItemCount := config.TotalUniqueLineItems()
		originalUnitCount := config.TotalUnits()
		originalAmount := config.TotalAmount()

		config.AddLineItem(lineItemA)

		Context("the config", func() {
			It("has 1 more unique line item than it had earlier", func() {
				Expect(config.TotalUniqueLineItems()).Should(Equal(originalLineItemCount + 1))
			})

			It("has 1 more unit than it had earlier", func() {
				Expect(config.TotalUnits()).Should(Equal(originalUnitCount + 1))
			})

			Specify("total amount increases by line item price", func() {
				Expect(config.TotalAmount()).Should(Equal(originalAmount + lineItemA.Price))
			})
		})
	})

	Context("when an existing line item is added", func() {
		config := Config{}

		config.AddLineItem(lineItemA)

		originalLineItemCount := config.TotalUniqueLineItems()
		originalUnitCount := config.TotalUnits()
		originalAmount := config.TotalAmount()

		config.AddLineItem(lineItemA)

		Context("the config", func() {
			It("has the same number of unique line items as earlier", func() {
				Expect(config.TotalUniqueLineItems()).Should(Equal(originalLineItemCount))
			})

			It("has 1 more unit than it had earlier", func() {
				Expect(config.TotalUnits()).Should(Equal(originalUnitCount + 1))
			})

			Specify("total amount increases by line item price", func() {
				Expect(config.TotalAmount()).Should(Equal(originalAmount + lineItemA.Price))
			})
		})
	})

	Context("that has 0 unit of line item A", func() {
		config := Config{}

		config.AddLineItem(lineItemB) // just to mimic the existence other line items
		config.AddLineItem(lineItemB) // just to mimic the existence other line items

		originalLineItemCount := config.TotalUniqueLineItems()
		originalUnitCount := config.TotalUnits()
		originalAmount := config.TotalAmount()

		Context("removing line item A", func() {
			config.RemoveLineItem(lineItemA.ID, 1)

			It("should not change the number of line items", func() {
				Expect(config.TotalUniqueLineItems()).Should(Equal(originalLineItemCount))
			})
			It("should not change the number of units", func() {
				Expect(config.TotalUnits()).Should(Equal(originalUnitCount))
			})
			It("should not change the amount", func() {
				Expect(config.TotalAmount()).Should(Equal(originalAmount))
			})
		})
	})

	Context("that has 1 unit of line item A", func() {
		config := Config{}

		config.AddLineItem(lineItemB) // just to mimic the existence other line items
		config.AddLineItem(lineItemB) // just to mimic the existence other line items

		config.AddLineItem(lineItemA)

		originalLineItemCount := config.TotalUniqueLineItems()
		originalUnitCount := config.TotalUnits()
		originalAmount := config.TotalAmount()

		Context("removing 1 unit line item A", func() {
			config.RemoveLineItem(lineItemA.ID, 1)

			It("should reduce the number of line items by 1", func() {
				Expect(config.TotalUniqueLineItems()).Should(Equal(originalLineItemCount - 1))
			})

			It("should reduce the number of units by 1", func() {
				Expect(config.TotalUnits()).Should(Equal(originalUnitCount - 1))
			})

			It("should reduce the amount by line item price", func() {
				Expect(config.TotalAmount()).Should(Equal(originalAmount - lineItemA.Price))
			})
		})
	})

	Context("that has 2 units of line item A", func() {

		Context("removing 1 unit of line item A", func() {
			config := Config{}

			config.AddLineItem(lineItemB) // just to mimic the existence other line items
			config.AddLineItem(lineItemB) // just to mimic the existence other line items
			//Reset the config with 2 units of line item A
			config.AddLineItem(lineItemA)
			config.AddLineItem(lineItemA)

			originalLineItemCount := config.TotalUniqueLineItems()
			originalUnitCount := config.TotalUnits()
			originalAmount := config.TotalAmount()

			config.RemoveLineItem(lineItemA.ID, 1)

			It("should not reduce the number of line items", func() {
				Expect(config.TotalUniqueLineItems()).Should(Equal(originalLineItemCount))
			})

			It("should reduce the number of units by 1", func() {
				Expect(config.TotalUnits()).Should(Equal(originalUnitCount - 1))
			})

			It("should reduce the amount by the line item price", func() {
				Expect(config.TotalAmount()).Should(Equal(originalAmount - lineItemA.Price))
			})
		})

		Context("removing 2 units of line item A", func() {
			config := Config{}

			config.AddLineItem(lineItemB) // just to mimic the existence other line items
			config.AddLineItem(lineItemB) // just to mimic the existence other line items
			//Reset the config with 2 units of line item A
			config.AddLineItem(lineItemA)
			config.AddLineItem(lineItemA)

			originalLineItemCount := config.TotalUniqueLineItems()
			originalUnitCount := config.TotalUnits()
			originalAmount := config.TotalAmount()

			config.RemoveLineItem(lineItemA.ID, 2)

			It("should reduce the number of line items by 1", func() {
				Expect(config.TotalUniqueLineItems()).Should(Equal(originalLineItemCount - 1))
			})

			It("should reduce the number of units by 2", func() {
				Expect(config.TotalUnits()).Should(Equal(originalUnitCount - 2))
			})

			It("should reduce the amount by twice the item price", func() {
				Expect(config.TotalAmount()).Should(Equal(originalAmount - 2*lineItemA.Price))
			})
		})

	})
})
