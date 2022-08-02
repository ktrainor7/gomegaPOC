package request

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/ghttp"
)

func TestBidRequest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Suite")
}

var _ = Describe("Client", func() {
	var (
		server     *ghttp.Server
		statusCode int
		body       []byte
		path       string
		addr       string
	)

	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("Test to local kraken server", func() {
		BeforeEach(func() {
			file, err := os.Create("data.txt")
			Expect(err).NotTo(HaveOccurred())
			body = []byte("Hi there!")
			file.Write(body)
			statusCode = 200
			path = "/read"
			addr = "http://" + server.Addr() + path

			BidRequestPayload := "%7B%22timeout%22%3A3000%2C%22cpmGranularity%22%3A1%2C%22cpmRange%22%3A%7B%22floor%22%3A0%2C%22ceil%22%3A20%7D%2C%22adSlotIds%22%3A%5B%22_adSlot3%22%5D%2C%22prebidRawBidRequests%22%3A%5B%7B%22bidder%22%3A%22kargo%22%2C%22params%22%3A%7B%22placementId%22%3A%22_jK5-oqWWnao%22%7D%2C%22auctionId%22%3A%22cad6d972-3aef-42de-bd95-4a94d21a798b%22%2C%22floorData%22%3A%7B%22skipped%22%3Afalse%2C%22skipRate%22%3A0%2C%22floorMin%22%3A5%2C%22location%22%3A%22setConfig%22%7D%2C%22userId%22%3A%7B%22pubcid%22%3A%22574e547f-7d76-43b0-a43b-ea608f3d99d1%22%7D%2C%22userIdAsEids%22%3A%5B%7B%22source%22%3A%22pubcid.org%22%2C%22uids%22%3A%5B%7B%22id%22%3A%22574e547f-7d76-43b0-a43b-ea608f3d99d1%22%2C%22atype%22%3A1%7D%5D%7D%5D%2C%22nativeParams%22%3A%7B%22sendTargetingKeys%22%3Afalse%2C%22image%22%3A%7B%22required%22%3Atrue%2C%22aspect_ratios%22%3A%5B%7B%22ratio_width%22%3A4%2C%22ratio_height%22%3A3%2C%22min_width%22%3A400%2C%22min_height%22%3A300%7D%5D%7D%2C%22title%22%3A%7B%22required%22%3Atrue%2C%22len%22%3A140%7D%2C%22sponsoredBy%22%3A%7B%22required%22%3Atrue%7D%2C%22clickUrl%22%3A%7B%22required%22%3Atrue%7D%2C%22body%22%3A%7B%22required%22%3Afalse%7D%7D%2C%22ortb2Imp%22%3A%7B%22ext%22%3A%7B%22data%22%3A%7B%22pbadslot%22%3A%22%2F8264%2Fmaw-cbsnews%2Fuplift%2Fmobile-mpu-plus-outstream-inc1%22%7D%7D%7D%2C%22mediaTypes%22%3A%7B%22banner%22%3A%7B%22sizes%22%3A%5B%5B300%2C250%5D%2C%5B320%2C50%5D%2C%5B11%2C11%5D%2C%5B1%2C1%5D%5D%7D%2C%22native%22%3A%7B%22sendTargetingKeys%22%3Afalse%2C%22image%22%3A%7B%22required%22%3Atrue%2C%22aspect_ratios%22%3A%5B%7B%22ratio_width%22%3A4%2C%22ratio_height%22%3A3%2C%22min_width%22%3A400%2C%22min_height%22%3A300%7D%5D%7D%2C%22title%22%3A%7B%22required%22%3Atrue%2C%22len%22%3A140%7D%2C%22sponsoredBy%22%3A%7B%22required%22%3Atrue%7D%2C%22clickUrl%22%3A%7B%22required%22%3Atrue%7D%2C%22body%22%3A%7B%22required%22%3Afalse%7D%7D%7D%2C%22adUnitCode%22%3A%22mobile-mpu-plus-outstream-inc1%22%2C%22transactionId%22%3A%22b62e1fcd-535c-4546-b196-26a96c90d021%22%2C%22sizes%22%3A%5B%5B300%2C250%5D%2C%5B320%2C50%5D%2C%5B11%2C11%5D%2C%5B1%2C1%5D%5D%2C%22bidId%22%3A%22751f14eb0fcef12%22%2C%22bidderRequestId%22%3A%22745f37d67fae8ed%22%2C%22src%22%3A%22client%22%2C%22bidRequestsCount%22%3A1%2C%22bidderRequestsCount%22%3A1%2C%22bidderWinsCount%22%3A0%7D%5D%2C%22userIDs%22%3A%7B%22kargoID%22%3A%22f390f11c-33b2-608e-afdf-4551e955a6e1%22%2C%22clientID%22%3A%223f45ded6-f9bb-49f5-b8c4-03067fcee0de%22%2C%22kargoCookieID%22%3A%220a0ab8f4-2976-4496-9507-ff6f4d2eea6f%22%2C%22crbIDs%22%3A%7B%222%22%3A%22d8a7875c-d9d7-4421-9f1a-99db89bdd654%22%2C%2216%22%3A%22X3TIeNHM5UkAABNnrrwAAAEu%26374%22%2C%2223%22%3A%22a8395f74-c87c-4700-aa9d-0a6f4bce9586%22%2C%2224%22%3A%22X3TIeNHM5UkAABNnrrwAAAEu%26374%22%2C%2225%22%3A%226fe02a6f-7ecc-4b95-95b5-36797581e614%22%2C%2229%22%3A%227533217711237367402%22%2C%2274%22%3A%22CAESENwF-x4BX6OaZEk2lHulCCg%22%2C%2282%22%3A%22GH8njz5ZO0iN%22%2C%2283%22%3A%22XaNKxsJjlbeBySP_E-RQPfJ5z3If8WGnELeYPql19TM%3D%22%2C%2287%22%3A%22AQEHTTzYpXSizQEIOlj4AQEBAQE%22%2C%2288%22%3A%22HGQCriui8lQ%22%2C%2290%22%3A%22dM_NoJwnQshiPmwiU1euamFXRjI%22%2C%2294%22%3A%22XY1612LZwat161w1oDPn6Zj6gTNm1-uG9big4dR-dik8mUd9g%22%2C%222_16%22%3A%22CAESENwF-x4BX6OaZEk2lHulCCg%22%2C%222_80%22%3A%22a8395f74-c87c-4700-aa9d-0a6f4bce9586%22%2C%222_93%22%3A%226fe02a6f-7ecc-4b95-95b5-36797581e614%22%2C%22Liveramp%22%3A%22XY1469IwwYrQn3Cqd0nTTpfzH76XJbgnoi86iopZFHlqHy6uo%22%2C%22Neustar%22%3A%22262070003560007916741%22%2C%22Tapad%22%3A%22452b0260-0347-11eb-80a0-7667e3ea56c8%22%2C%22ttdo%22%3A%22c073c8ff-ff7e-4c1c-9359-31cc1bc65c48%22%7D%2C%22optOut%22%3Afalse%7D%2C%22pageURL%22%3A%22https%3A%2F%2Fwww.newstimes.com%2Fnews%2Farticle%2FCT-woman-who-died-from-COVID-remembered-for-her-16947901.php%3Ft%3D019c8bcbc3%26utm_source%3Dnewsletter%26utm_medium%3Demail%26utm_campaign%3DCT_NT_MorningBriefing%26sid%3D591c8fd624c17c3e4b8c327b%22%7D"

			BidrequestURL := fmt.Sprintf("https://kraken-master-kraken.staging.kargo.com/api/v1/bid?json=%s", url.QueryEscape(BidRequestPayload))

			ghttp.VerifyRequest("GET", BidrequestURL)
			bdy, err := getResponse(BidrequestURL)

			fmt.Sprintf(string(bdy))
			Expect(err).ShouldNot(HaveOccurred())

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})
		AfterEach(func() {
			err := os.Remove("data.txt")
			Expect(err).NotTo(HaveOccurred())
		})
		It("Reads data from file successfully", func() {
			bdy, err := getResponse(addr)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bdy).To(Equal(body))
		})
	})
})
