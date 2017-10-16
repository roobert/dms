package client

//import (
//	"time"
//
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//)
//
//var _ = Describe("Client", func() {
//	t := time.Now()
//	c := Client{Name: "test", TimeStamp: t}
//
//	Describe("Converting timestmap to date", func() {
//		Context("With the timestamp", func() {
//			It("should convert to a date string", func() {
//				Expect(c.Date()).To(Equal(t.Format("2006-01-02")))
//			})
//		})
//	})
//
//	Describe("Converting timestmap to slot", func() {
//		Context("With the timestamp", func() {
//			//It("should convert to a slot number", func() {
//			//	//Expect(c.Slot()).To(Equal(int64))
//			//	Ω(c.Slot()).Should(BeAssignableToTypeOf(int64))
//			//})
//			It("should be within range: 0..5759", func() {
//				Expect(c.Slot() > int64(0) && c.Slot() < int64(5799)).To(BeTrue())
//			})
//		})
//	})
//
//})
