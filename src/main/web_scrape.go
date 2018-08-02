package main

import (
	"fmt"
	"../web"
)

func main() {
	fmt.Println("START")

	//var t web.TargetPage = web.NewYahoo()
	t := web.NewYahoo()
	t.GetPage("origin=http%3A%2F%2Fheadlines.yahoo.co.jp&sort=lost_points&order=desc&page=1&type=t&keys=WMznMmODKGC5MRxxPMOoZdqWbnNKsKvnchyUt_J2pRnWKPYas7YNfolJSpOCYrlmFwTohxGJ_w2A8DR27w7b3BZujqe6JoH4x5V._N.k41gQPSoEbchomEQ_QvLnnY_97Aw.P7avu5pIiKppjR7wwEZa67VNIGsEVWLl7VWqhobeuuAtlrKfboOONRjC7MohYfGjubKsDneui4gL23Ux3E1qLRKncJdk2uZN09sdAulNUQpQVYWx7vCEXzdg_YNpLy1jcOUZ2TMC2l6_d9wiK7OXgS67A1jjxCV5l96O5ebowpXLrTo2AEdjXsA_PxI_HUOlPn1hF6UHwkbeSikz2HamkHHoSOLAgC.ogulem5BKLZdFDeKcn3RSfyVSdLm.cnPzsBwswZsC9ZOlCuRD7SbWpZVxdS3x64z91MdWjkxE53xePf4h3x5WMu6tojqqA342_9x.KmpFoooEKDeWtYmFqOEOuwg9rr1Q_7qBZk9BRCwl.ZX5T05E_5G47SypVNsMy_wyfVv.TyxlYMjOTskF6POdB87oghNf7dQN4rtrgEXYE_K3V5RxhuETCgeUiFOyiLB.qxT8xEkHUJjKqWuelNw4xtAQHyDE0VFXFzLd3Fete54gQylX00Hnng4BNRKt3JDSthQS3CtppfBPqomJ6NAr4zs9REGTk332Guw5UW1xV6RFQRdSnBxXWRAwUnHmeDV5cXi.Wbm06NRIbOpPmrSWpCKk1j357UqZPVOVwieHZT9yJPtpWbbda.OX0fQ.qemoOanWawApsLqy&full_page_url=http%3A%2F%2Fheadlines.yahoo.co.jp%2Fcm%2Fmain%3Fd%3D20170118-00000034-zdn_mkt-bus_all&comment_num=10&ref=&bkt=&flt=&disable_total_count=&ua=Mozilla%2F5.0+(Macintosh%3B+Intel+Mac+OS+X+10.11%3B+rv%3A52.0)+Gecko%2F20100101+Firefox%2F52.0")

	fmt.Println("******************************************************")

	t2 := web.NewRuiji()
	t2.GetPage("title=%E4%B8%8D%E8%83%BD%E7%8A%AF")

	fmt.Println("END")
}
