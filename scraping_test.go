package main

import (
	"testing"
)

func Test_parseDonatePageURL(t *testing.T) {
	type args struct {
		contents string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "parses correctly from snapshot of page from 2024-06-02",
			args: args{
				contents: `      <!DOCTYPE html>
				<!--[if lt IE 7]>
				<html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
				<!--[if IE 7]>
				<html class="no-js lt-ie9 lt-ie8"> <![endif]-->
				<!--[if IE 8]>
				<html class="no-js lt-ie9"> <![endif]-->
				<!--[if gt IE 8]><!-->
				<html class="no-js"> <!--<![endif]-->
				<head>
				  <meta name="google-site-verification" content="pJgtWPNLX5_ucuNccfFyJu8LSQ-aDzrnVveSiUc23x8" />
				  <meta name="google-site-verification" content="CjNDrfIadB9PJb8DiRI7IDIC2Ud_s4OdC1A9eLfNXP8" />
				  <meta name="google-site-verification" content="tWVq_LM8h4MFZ51wf5z6L8SzXyA115_g_ie4MFOKijg" />
				  <meta charset="utf-8">
				  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
				  <title>Charter Challenge for Fair Voting</title>
				  <meta name="viewport" content="width=device-width,initial-scale=1">
				  <link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css" rel="stylesheet">
				  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/theme.scss" type="text/css" />
				  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/tablet-and-desktop.scss" type="text/css" media="screen and (min-width: 768px)" />

				  <!-- because ie8 ignores media queries, we need this -->
				  <!--[if IE 8]>
					<link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/tablet-and-desktop.scss" type="text/css" />
				  <![endif]-->


				  <!--[if IE]>
				  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/ie.scss" type="text/css" />
				  <![endif]-->



				<script type="text/javascript">var _sf_startpt=(new Date()).getTime()</script>



				<meta content="authenticity_token" name="csrf-param" />
				<meta content="vRw8PxR5c7Xb0QsfmIB9TVLL3SWNVx/ZK2sugu3oukI=" name="csrf-token" />

				  <link rel="canonical" href="https://www.charterchallenge.ca/" />
					<meta name="Title" content="About - Charter Challenge for Fair Voting">
					<meta name="Description" content="Since the 1980s Canadian courts have been ruling in favour of fair and equal representation in Canadian and provincial elections. Now, it&#39;s time to challenge the discriminatory first-past-the-post voting system itself.
				">
					<meta property="og:title" content="Charter Challenge for Fair Voting"/>
				  <meta property="og:url" content="https://www.charterchallenge.ca/">
					<meta property="og:description" content="Since the 1980s Canadian courts have been ruling in favour of fair and equal representation in Canadian and provincial elections. Now, it&#39;s time to challenge the discriminatory first-past-the-post voting system itself.
				">
					<meta property="og:type" content="article">
					  <link rel="image_src" href="https://assets.nationbuilder.com/springtide/pages/369/meta_images/original/Charter_Challenge_LOGO-min.png?1497226761" />
					  <meta property="og:image" content="https://assets.nationbuilder.com/springtide/pages/369/meta_images/original/Charter_Challenge_LOGO-min.png?1497226761" />
				  <meta property="og:site_name" content="Charter Challenge for Fair Voting"/>

				<script type="text/javascript">
				  var NB = NB || {};

				  NB.environment = "production";
				  NB.pageId = "369";
				  NB.Liquid = NB.Liquid || {
					Theme: {
					  version: 2,
					  name: "Charter Challenge for Fair Voting",
					  parent: "publish",
					  variation: ""
					}
				  };
				  NB.payments = NB.payments || {elements: {}};
				  NB.payments.decimal_mark = "."
				  NB.payments.thousands_separator = ","
				  NB.payments.currency = "cad";
				</script>

				<script type="text/javascript">
					//<![CDATA[
					  window._auth_token_name = "authenticity_token";
					  window._auth_token = "vRw8PxR5c7Xb0QsfmIB9TVLL3SWNVx/ZK2sugu3oukI=";
					//]]>
				</script>

				  <link rel="stylesheet" href="https://ajax.googleapis.com/ajax/libs/jqueryui/1.10.0/themes/cupertino/jquery-ui.css" type="text/css" media="all">

					  <link rel="icon" type="image/x-icon" href="https://assets.nationbuilder.com/springtide/sites/38/favicon_images/original/faviconcc.png?1496167237" />


				  <script src="https://assets.nationbuilder.com/assets/liquid/main-f52182358767f5af49bb34ddeeedb502a15f0105ffc14c758599b2789870803b.js"></script>

				<script type="text/javascript">
				  window.twttr = (function (d,s,id) {
					var t, js, fjs = d.getElementsByTagName(s)[0];
					if (d.getElementById(id)) return; js=d.createElement(s); js.id=id;
					js.src="//platform.twitter.com/widgets.js"; fjs.parentNode.insertBefore(js, fjs);
					return window.twttr || (t = { _e: [], ready: function(f){ t._e.push(f) } });
				  }(document, "script", "twitter-wjs"));
				</script>

				<script type="text/javascript">
				  NB.FBAppId = '1679673749184434';
				</script>

				  <script type="text/javascript">
					(function (d) { var config = { kitId: 'atg5ome', scriptTimeout: 3000, async: true }, h=d.documentElement,t=setTimeout(function(){h.className=h.className.replace(/\bwf-loading\b/g,"")+" wf-inactive";},config.scriptTimeout),tk=d.createElement("script"),f=false,s=d.getElementsByTagName("script")[0],a;h.className+=" wf-loading";tk.src='https://use.typekit.net/'+config.kitId+'.js';tk.async=true;tk.onload=tk.onreadystatechange=function(){a=this.readyState;if(f||a&&a!="complete"&&a!="loaded")return;f=true;clearTimeout(t);try{Typekit.load(config)}catch(e){}};s.parentNode.insertBefore(tk,s)})(document);
				  </script>

				  <script type="text/javascript">
					var _gaq = _gaq || [];
					_gaq.push(['_setAccount', 'UA-100048182-1']);
					 _gaq.push(['_setDomainName', 'none']);
					_gaq.push(['_gat._anonymizeIp']);
					_gaq.push(['_setAllowLinker', true]);
					  _gaq.push(['_setCustomVar', 1, 'UGC', 'false', 3]);
					  _gaq.push(['_setCustomVar', 1, 'Page type', 'Basic', 3]);
					_gaq.push(['_trackPageview']);
					_gaq.push(['_trackPageLoadTime']);

					(function() {
					  var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
						  ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';

					  let s = document.getElementsByTagName('script')[0];
					  s.parentNode.insertBefore(ga, s);
					})();
				  </script>








				<!-- The following line of CSS hides a checkbox in social share prompts for posting to Facebook;
				As of Aug 1 2018, FB has deprecated the ability for apps to post to personal profile pages, so this is a temporary fix
				while we re-configure the social share prompt -->
				<style>label[for="face_tweet_is_facebook"]{display:none;}</style>



				  <script type="text/javascript">
					NB.appConfig.userIsLoggedIn = false;
				  </script>

				  <link href='//fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic,700italic|Oswald:400,300,700|Lato:400,700,400italic,700italic|Bree+Serif|Josefin+Sans:700|Bitter:400,700,400italic|Open+Sans:400italic,700italic,400,700,800' rel='stylesheet' type='text/css'>
				<link href="https://fonts.googleapis.com/css?family=Open+Sans:400,400i,700,700i" rel="stylesheet">
				<link href="https://fonts.googleapis.com/css?family=Lato:900" rel="stylesheet">
				<!-- Facebook Pixel Code -->
				<script>
				  !function(f,b,e,v,n,t,s)
				  {if(f.fbq)return;n=f.fbq=function(){n.callMethod?
				  n.callMethod.apply(n,arguments):n.queue.push(arguments)};
				  if(!f._fbq)f._fbq=n;n.push=n;n.loaded=!0;n.version='2.0';
				  n.queue=[];t=b.createElement(e);t.async=!0;
				  t.src=v;s=b.getElementsByTagName(e)[0];
				  s.parentNode.insertBefore(t,s)}(window, document,'script',
				  'https://connect.facebook.net/en_US/fbevents.js');
				  fbq('init', '119846658737416');
				  fbq('track', 'PageView');
				</script>
				<noscript><img height="1" width="1" style="display:none" src="https://www.facebook.com/tr?id=119846658737416&ev=PageView&noscript=1" /></noscript>
				<!-- End Facebook Pixel Code -->

				<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/jquery.ui.effect.min.js"></script>
				<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/jquery.ui.effect-slide.min.js"></script>
				<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/staged-donations.js"></script>

				</head>
				<body class="aware-theme v2-theme page-type-basic js ">
				 <!-- Google Tag Manager (noscript) -->
				<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-P6TB763" height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
				<!-- End Google Tag Manager (noscript) -->


				  <div id="pattern" class="pattern">
					<div class="wrap" id="wrap">
					  <div id="body" class="page-pages-show-basic">

						<div class="header-container clearfix">
						  <div class="width-container clearfix">

							<div class="tablet-visible">

				<!-- _nav.html -->

				<nav id="menu" role="navigation">
				  <ul id="topnav" class="topnav desktop-nav">


					<li class='active '>
					  <a href="/">Home</a>
					</li>



					<li class=''>
					  <a href="/why">Why a challenge? </a>
					</li>



					<li class='drop'>
					  <a href="/the_plan">The Plan <i class="icon-angle-down"></i></a>
					  <ul class="sub">

						<li class="show-parent"><a href="/the_plan">The Plan</a></li>


						<li><a href="/case_evidence">Case Evidence </a></li>

						<li><a href="/ruling">Ruling</a></li>

						<li><a href="/appeal">Appeal</a></li>

					  </ul>
					</li>



					<li class=''>
					  <a href="/faq">FAQ</a>
					</li>



					<li class=''>
					  <a href="/contact">Contact</a>
					</li>



					<li class=''>
					  <a href="/blog">Blog Updates</a>
					</li>



					<li class=''>
					  <a href="/team">Team</a>
					</li>



					<li class=''>
					  <a href="/signup_list">Get Updates</a>
					</li>



					<li class=''>
					  <a href="/donate_appeal">Donate</a>
					</li>


				  </ul>
				</nav>

				<a href="#menu" class="menu-link"><i class="icon-menu"></i></a>


				<!-- /_nav.html -->

							</div>


							  <div class="site-logo">
								<header><a href="/"><img src="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891"></a></header>
							  </div>



						  </div>
						  <!-- .width-container -->
						</div>
						<!-- .header-container -->

						<div class="nav-container desktop-visible">
						  <div class="width-container clearfix">

				<!-- _nav.html -->

				<nav id="menu" role="navigation">
				  <ul id="topnav" class="topnav desktop-nav">


					<li class='active '>
					  <a href="/">Home</a>
					</li>



					<li class=''>
					  <a href="/why">Why a challenge? </a>
					</li>



					<li class='drop'>
					  <a href="/the_plan">The Plan <i class="icon-angle-down"></i></a>
					  <ul class="sub">

						<li class="show-parent"><a href="/the_plan">The Plan</a></li>


						<li><a href="/case_evidence">Case Evidence </a></li>

						<li><a href="/ruling">Ruling</a></li>

						<li><a href="/appeal">Appeal</a></li>

					  </ul>
					</li>



					<li class=''>
					  <a href="/faq">FAQ</a>
					</li>



					<li class=''>
					  <a href="/contact">Contact</a>
					</li>



					<li class=''>
					  <a href="/blog">Blog Updates</a>
					</li>



					<li class=''>
					  <a href="/team">Team</a>
					</li>



					<li class=''>
					  <a href="/signup_list">Get Updates</a>
					</li>



					<li class=''>
					  <a href="/donate_appeal">Donate</a>
					</li>


				  </ul>
				</nav>

				<a href="#menu" class="menu-link"><i class="icon-menu"></i></a>


				<!-- /_nav.html -->

						  </div>
						</div>



						  <div class="width-container">

				<!-- _features.html -->







				  <div id="page-features" class="page-features">


				  </div>





				<!-- /_features.html -->

						  </div>


						<div class="main-container" id="middle">
						  <div class="main width-container clearfix">



				<!-- _columns_2.html -->
				<div class="twocolumn-container clearfix">

				  <div class="left-column">

					<div id="flash_container">



				</div>


					<div class="content-pages-show-basic">

				<!-- _breadcrumbs.html -->




				<!-- /_breadcrumbs.html -->




				<div id="content">

				  <div id="intro" class="intro">
					<div class="text-content">
					  <p><img src="https://d3n8a8pro7vhmx.cloudfront.net/springtide/pages/369/attachments/original/1559228052/Charter_Challenge_-_FYI_POST_1_-_OUR_RIGHTS__FACEBOOK.png?1559228052" alt="When politicians don't stand up for our civil rights, it's up to citizens to fight for those rights in court. Nobody else can do it for us. " width="1200" height="627"></p>
				<p>Since the Canadian Charter of rights and freedoms was established in the 1980s, the courts have been ruling in favour of fair and equal representation in Canadian and provincial elections. Now, we're challenging the (un)fairness of the first-past-the-post voting system itself. </p>
				<p><a href="https://www.charterchallenge.ca/donate_appeal">Donate</a> to make sure our challenge is fully heard in court. </p>
					</div>
				  </div>




					<ul class="homepage_excerpt-list">

				  <li class="excerpt-block"><hr>
				<!-- _homepage_excerpt.html -->
				<div class="homepage-excerpt page-379">



					<i class="widget-icon fa fa-comments-o"></i>
					<h3 class="excerpt-type"><a href="/blog">Latest from the blog</a></h3>

					<div class="row-fluid blog-widget-parent">



					  <div class="span6 blog-widget">



						<div class="blog-widget-content">
						  <div class="byline">

							<div class="byline-image-wrap">
							  <div class="byline-image">
								<img src="https://assets.nationbuilder.com/springtide/profile_images/8c8dc0d16db90ce692cf534792dbda8d31378508.jpg?1603294239" border="0" width="72" height="72" class="profile_image" title="Jesse Hitchcock" alt="Jesse Hitchcock"/>
							  </div>


							</div>

							posted by <span class="name"><span class="linked-signup-name">Jesse Hitchcock</span></span> | <span class="pc">5sc</span><br>
							  August 18, 2023


						  </div>

						  <h4 class="excerpt-title"><a href="/do_we_have_to_waste_half_the_votes">Do We Have to Waste Half the Votes?</a></h4>


						  <div class="truncate-300" data-truncate="300">
						  As we shared with you in our last message, we’ve now submitted all our written arguments for the upcoming Charter Challenge court hearing (Sep 25-27 in Toronto).  We’re currently waiting to receive the government’s written argument and are now working on our oral arguments.
				In the meantime, we wanted to share with you a few more details from Fair Voting BC’s affidavit, which presents some insightful metrics we’re using to reinforce a number of our key arguments.
				In a previous message, we reported that over half the voters in the 2019 election cast “wasted votes” (votes that didn’t help elect an MP). While this was true in that election, the court will want to know if this is a consistent pattern across time, and if changing the voting system will likely change this.
				To answer this question, we defined a Representation Metric that could be calculated both for countries such as Canada that use the First Past the Post voting system and for countries that use proportional voting.
				More Than Half the Votes Are Wasted in Canadian Elections
				In Canada, the Representation Metric is simply the proportion of voters who cast a vote for an elected candidate. Thanks to work by Byron Weber Becker (University of Waterloo), we were able to show that for nearly the past hundred years (since 1935) the Representation Metric has averaged less than 50% over the past 40 years.

				Representation Metric for all federal elections in Canada since Confederation. The Representation Metric expresses the percentage of voters who have voted for an elected MP. Since 1984, the Representation Metric has averaged just 49.8%.
				Are As Many Votes Wasted Under Proportional Voting?
				Absolutely not.  In Proportional Representation (PR) voting systems, the Representation Metric is typically much higher (and the “wasted votes” much lower), because far higher proportions of the ballots cast affect the resulting makeup of the legislature. To demonstrate this point, we analyzed the results of recent elections in three countries using different types of PR systems (Norway (List PR), New Zealand (MMP), and Ireland (STV)).  Under all three of these systems, the Representation Metric would have been far higher, ranging from about 77-79% under Ireland’s ‘direct representation’ STV system up to as high as 94-95% under Norway’s List Proportional Representation system.

				Representation Metrics for Canada (2019 federal election) and two recent elections in each of three selected countries using proportional voting (to demonstrate consistency of findings). The darker shade represents those voters who are represented by a representative for whom they have explicitly voted (‘direct’ representation), while the lighter shade represents those voters who are represented by a representative who won a seat by virtue of votes cast for the representative’s party (‘indirect’ representation).
				This evidence is critical for us to make the points that (1) our recent experience in Canada is not an anomaly, but rather a routine and expected outcome of elections under FPTP, and (2) any number of alternative voting systems would markedly increase the proportion of represented voters.
				Charter Challenge In the News
				The Charter Challenge is in the news! Maxwell Cameron, a professor in the Department of Political Science and the School of Public Policy and Global Affairs at the University of British Columbia, contributed an opinion piece to The Globe and Mail titled “The Charter challenge of first-past-the-post could lead to a better electoral system.” To read the piece, click here.
				Thank you so much for your continued support. More details shortly!
				Jesse Hitchcock, Springtide &amp; Antony Hodgson, Fair Voting BC</div>

						  <span class="read-more"><a href="/do_we_have_to_waste_half_the_votes">read more</a></span>
						</div>

						<div class="share-bar">


						  <a href="/do_we_have_to_waste_half_the_votes#addreaction" class="button icon-button">

							  <i>1</i> comment

						  </a>



						  <a href="/forms/shares/new?page_id=1777" class="get button icon-button"><i class="fa fa-sitemap"></i>Share</a>
						  <div id="page-share-form-1777"></div>


						</div>

					  </div>



					  <div class="span6 blog-widget">



						<div class="blog-widget-content">
						  <div class="byline">

							<div class="byline-image-wrap">
							  <div class="byline-image">
								<img src="https://assets.nationbuilder.com/springtide/profile_images/8c8dc0d16db90ce692cf534792dbda8d31378508.jpg?1603294239" border="0" width="72" height="72" class="profile_image" title="Jesse Hitchcock" alt="Jesse Hitchcock"/>
							  </div>


							</div>

							posted by <span class="name"><span class="linked-signup-name">Jesse Hitchcock</span></span> | <span class="pc">5sc</span><br>
							  August 18, 2023


						  </div>

						  <h4 class="excerpt-title"><a href="/charter_challenge_update_in_court_next_month">Charter Challenge Update - In Court Next Month!</a></h4>


						  <div class="truncate-300" data-truncate="300">
						  Our date in court is approaching rapidly and there is important and exciting work underway.
				The Charter Challenge for Fair Voting will be heard in the Ontario Superior Court in Toronto next month (Sep 25-27). The hearings will be open to the public, and we’ll send you all the details as soon as we have them confirmed.
				In the meantime, here’s an update on the work over the past month and a half. At the end of June, we told you that, thanks to your generous support, we had reached the very significant milestone of having submitted our factum (the written argument for the case).
				Since then, four external organizations have submitted supplementary arguments of their own (called interventions) - three supporting our case, and one arguing against.


				Fair Vote Canada submitted a review of electoral reform attempts throughout Canadian history, clearly demonstrating that previous efforts were thwarted by partisan self-interest and urging the court to act to defend voters’ democratic rights.


				The UK-based Electoral Reform Society reinforced many of our core arguments, and noted that the issues we identified play out with similar effects in the UK.


				Apathy is Boring is a Canadian advocacy organization seeking to improve engagement and participation of youth in our democracy. Their submission argues that our current voting system creates a number of significant barriers to participation by and representation of younger voters.


				Finally, the Canadian Constitution Foundation doesn’t directly take a position on the relative merits of different voting systems, but argues that our first-past-the-post system is “constitutionalized” and that the court cannot force the government to change it. We disagree, and our lawyer, Nicolas Rouleau, submitted a compelling and well-argued response to their submission.


				What’s Next?
				The federal government is expected to submit their written argument very shortly, at which point all the written documents will be released publicly. The team will read the government’s argument very carefully, and take it into account as we prepare oral arguments.
				Thank you for your continued support. More to come soon!
				Jesse Hitchcock, Springtide &amp; Antony Hodgson, Fair Voting BC</div>

						  <span class="read-more"><a href="/charter_challenge_update_in_court_next_month">read more</a></span>
						</div>

						<div class="share-bar">


						  <a href="/charter_challenge_update_in_court_next_month#addreaction" class="button icon-button">

							  <i>1</i> comment

						  </a>



						  <a href="/forms/shares/new?page_id=1776" class="get button icon-button"><i class="fa fa-sitemap"></i>Share</a>
						  <div id="page-share-form-1776"></div>


						</div>

					  </div>



					</div>


					<span class="continue">

					  <a class="button submit-button" href="/blog">See all posts</a>


					</span>

				  <!-- Blog post page -->


				</div> <!-- .homepage-excerpt -->
				<!-- /_homepage_excerpt.html -->
				</li>

				</ul>


				  <div class="like-page">
					<strong>Do you like this page?</strong>

				<!-- _like_page.html -->
				<div class="sharetable padtopless clearfix">
				  <div class="facebook-cell">

					<div class="desktop-visible">
					  <fb:like href="https://www.charterchallenge.ca/" show_faces="false" width="300" colorscheme="light" send="true" action="like"></fb:like>
					</div>
					<div class="mobile-visible">
					  <fb:like href="https://www.charterchallenge.ca/" show_faces="false" width="245" colorscheme="light" send="false" action="like"></fb:like>
					</div>

				  </div>
				  <div class="twitter-cell">

					  <a href="https://twitter.com/share" class="twitter-share-button" data-count="none" data-text="It's time to challenge Canada's discriminatory first-past-the-post voting system. " data-via="SpringtideCo">Tweet</a>

				  </div>
				</div>

				<!-- /_like_page.html -->

				  </div>



				<!-- _page_stream.html -->


				<!-- /_page_stream.html -->



				</div>
					</div>

				  </div>
				  <!-- .left_column -->

				  <hr class="mobile-visible"/>

				  <div class="right-column">



					 <p></p>



				 <p>
					  <a href="https://www.charterchallenge.ca/donate_appeal">
				   <img src="https://d3n8a8pro7vhmx.cloudfront.net/themes/592ad4daed0e46ae54000000/attachments/original/1556450696/CC_-_WEB_BUTTONS_DONATE_-_RED.png?1556450696">
				   </a></p>
				   <p><br>
					<a href="/signup_list">
				   <img src="https://d3n8a8pro7vhmx.cloudfront.net/themes/592ad4daed0e46ae54000000/attachments/original/1556450696/CC_-_WEB_BUTTONS_UPDATES.png?1556450696">
					</a>
					</p>

					 <p>Like and Follow</p>
					  <ul class="supporter-nav-signin">
					   <li><a href="https://www.facebook.com/charterchallenge/"><i class="fa fa-facebook"></i></a>
						</li>
						<li><a href="https://twitter.com/Challenge4FV"><i class="fa fa-twitter"></i></a></li>
						<li><a href="https://www.instagram.com/charterchallenge"><i class="fa fa-instagram"></i></a></li>
					  </ul>
						  <p> -			 -			 - </p>

				<h3>Key Events in the Case:</h3>
				<ul>
					<li> We're currently waiting for our appeal to be heard (scheduled for Nov 2024)
					<li> We submitted our appeal factum in late April 2024.
					<li> We filed our Notice of Appeal on Dec 29, 2023.
					<li> Justice Ed Morgan issued his ruling on Nov 30, 2023 and unfortunately dismissed our application.
					<li> The case was heard September 26-28, 2023 in the Ontario Superior Court.
					<li> We received the government's affidavits in late Fall 2022.
					<li> We served the government with our affidavit and evidence package in May 2021.
					<li> We filed the case with the Ontario Superior Court in October 2019.
				</ul>

				<p><h3>How you can help</h3></p>

				  <p>The main way you can help is to support the case financially. We are now raising $25,000 to support the second stage of our appeal process. You can support the case for as little as a dollar a month.

				<p><h3>What to expect</h3></p>

				<p>At each step, we set a donation goal based on our estimate of the costs for the next stage of the process, and invite our supporters to contribute towards that goal to ensure the case can continue to move forward.</p>
				<br>

					<div class="signin-wrap">
						  <p> -			 -			 - </p>


				<p>This project is a collaboration between <a href="https://www.springtide.ngo">Springtide</a> and <a href="https://fairvotingbc.com">Fair Voting BC</a>.</p>



				  </div>
				  <!-- .right-column -->
				</div> <!-- .twocolumn_container -->

				<!-- /_columns_2.html -->



						  </div>
						  <!-- .main -->
						</div>
						<!-- .main-container -->

						<footer>
						  <div class="width-container clearfix">

							<div id="fb-root"></div>

				<script type="text/javascript">
				  window.fbAsyncInit = function() {
					FB.init({
					  appId      : 1679673749184434,
					  channelUrl : "//cc-springtide.nationbuilder.com/channel.html",
					  status     : true,
					  version    : "v18.0",
					  cookie     : true,
					  xfbml      : true
					});
				  };
				  (function(d, s, id) {
					var js, fjs = d.getElementsByTagName(s)[0];
					if (d.getElementById(id)) return;
					js = d.createElement(s);
					js.id = id;
					js.async = true;
					js.src = "//connect.facebook.net/en_US/sdk.js";
					fjs.parentNode.insertBefore(js, fjs);
				  }(document, 'script', 'facebook-jssdk'));
				</script>

				<script src="https://assets.nationbuilder.com/assets/liquid-6fda76e47cd1a46bec92e2adac0a0453c78638197e234d7667c2ff4366c5a44a.js"></script>








							<div class="row-fluid">

							  <div class="span6">

							  </div>



							</div>

						  </div>
						  <!-- .width-container -->
						</footer>

					  </div>
					  <!-- #body -->
					</div>
					<!-- #wrap -->
				  </div>
				  <!-- #pattern -->





				<script type="text/javascript">
				  function pollFBCommentBox(widget) {
					var fbCommentBoxTimer = window.setInterval(function() {
					  if($(widget).height() < 100) {
						$(widget).removeClass('comment-box-active');
						window.clearInterval(fbCommentBoxTimer);
					  }
					}, 1000);
				  }
				  FB.Event.subscribe('edge.create', function(href, widget) {
					$(widget).addClass('comment-box-active');
					pollFBCommentBox(widget);
				  });
				  FB.Event.subscribe('edge.remove', function(href, widget) {
					$(widget).removeClass('comment-box-active');
				  });
				</script>

				  <script>
					if ( window.self !== window.top ) {
					  var referrer_origin = "";
					  if ( window.location.origin !== referrer_origin ) {
						var xhttp = new XMLHttpRequest();
						var params = "iframe_req_path=" + window.location.pathname + "&referrer_origin=" + referrer_origin;
						xhttp.open("GET", "iframe_requests?" + params, true);
						xhttp.send();
					  }
					}
				  </script>
				<script defer src="https://static.cloudflareinsights.com/beacon.min.js/vef91dfe02fce4ee0ad053f6de4f175db1715022073587" integrity="sha512-sDIX0kl85v1Cl5tu4WGLZCpH/dV9OHbA4YlKCuCiMmOQIk4buzoYDZSFj+TvC71mOBLh8CDC/REgE0GX0xcbjA==" data-cf-beacon='{"rayId":"88da9fd18933aafd","version":"2024.4.1","token":"ff633690ee944cc4bb4fa2453466614e"}' crossorigin="anonymous"></script>
				</body>

				</html>


				`,
			},
			want:    "https://www.charterchallenge.ca/donate_appeal",
			wantErr: false,
		},
		{
			name: "returns error for page with no matching URL - 1",
			args: args{
				contents: `<html><a href="https://charterchallenge.ca/somepage></a></html>`,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "returns error for page with no matching URL - 2",
			args: args{
				contents: `<html><a href="https://example.com/donate></a></html>`,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDonatePageURL(tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDonatePageURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseDonatePageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseDonatedSoFar(t *testing.T) {
	type args struct {
		webPage string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Parses $5,526.00 correctly from the web page returned from my web browser on Jan 8, 2023.",
			args: args{
				webPage: webPageWith5526,
			},
			want:    "$5,526.00",
			wantErr: false,
		},
		{
			name: "Parses $0 correctly from the web page returned from my web browser on Sep 12, 2023, which displays the substring JUST STARTED instead of a monetary value.",
			args: args{
				webPage: webPageWithJustStarted,
			},
			want:    "$0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDonatedSoFar(tt.args.webPage)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDonatedSoFar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseDonatedSoFar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseGoal(t *testing.T) {
	type args struct {
		pageContents string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "parses correctly from snapshot of page from 2024-06-02",
			args: args{
				pageContents: `      <!DOCTYPE html>
				<!--[if lt IE 7]>
				<html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
				<!--[if IE 7]>
				<html class="no-js lt-ie9 lt-ie8"> <![endif]-->
				<!--[if IE 8]>
				<html class="no-js lt-ie9"> <![endif]-->
				<!--[if gt IE 8]><!-->
				<html class="no-js"> <!--<![endif]-->
				<head>
				  <meta name="google-site-verification" content="pJgtWPNLX5_ucuNccfFyJu8LSQ-aDzrnVveSiUc23x8" />
				  <meta name="google-site-verification" content="CjNDrfIadB9PJb8DiRI7IDIC2Ud_s4OdC1A9eLfNXP8" />
				  <meta name="google-site-verification" content="tWVq_LM8h4MFZ51wf5z6L8SzXyA115_g_ie4MFOKijg" />
				  <meta charset="utf-8">
				  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
				  <title>Donate Charter Challenge for Fair Voting</title>
				  <meta name="viewport" content="width=device-width,initial-scale=1">
				  <link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css" rel="stylesheet">
				  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/theme.scss" type="text/css" />
				  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/tablet-and-desktop.scss" type="text/css" media="screen and (min-width: 768px)" />

				  <!-- because ie8 ignores media queries, we need this -->
				  <!--[if IE 8]>
					<link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/tablet-and-desktop.scss" type="text/css" />
				  <![endif]-->


				  <!--[if IE]>
				  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/ie.scss" type="text/css" />
				  <![endif]-->



				<script type="text/javascript">var _sf_startpt=(new Date()).getTime()</script>



				<meta content="authenticity_token" name="csrf-param" />
				<meta content="vRw8PxR5c7Xb0QsfmIB9TVLL3SWNVx/ZK2sugu3oukI=" name="csrf-token" />

				  <link rel="canonical" href="https://www.charterchallenge.ca/donate_appeal" />
					<meta name="Title" content="Appeal Fund Stage 2 - Going to Court Again!">
					<meta property="og:title" content="Appeal Fund Stage 2 - Going to Court Again!"/>
				  <meta property="og:url" content="https://www.charterchallenge.ca/donate_appeal">
					<meta property="og:type" content="article">
					  <link rel="image_src" href="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891" />
					  <meta property="og:image" content="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891" />
				  <meta property="og:site_name" content="Charter Challenge for Fair Voting"/>

				<script type="text/javascript">
				  var NB = NB || {};

				  NB.environment = "production";
				  NB.pageId = "1780";
				  NB.Liquid = NB.Liquid || {
					Theme: {
					  version: 2,
					  name: "Charter Challenge for Fair Voting",
					  parent: "publish",
					  variation: ""
					}
				  };
				  NB.payments = NB.payments || {elements: {}};
				  NB.payments.decimal_mark = "."
				  NB.payments.thousands_separator = ","
				  NB.payments.currency = "cad";
				</script>

				<script type="text/javascript">
					//<![CDATA[
					  window._auth_token_name = "authenticity_token";
					  window._auth_token = "vRw8PxR5c7Xb0QsfmIB9TVLL3SWNVx/ZK2sugu3oukI=";
					//]]>
				</script>

				  <link rel="stylesheet" href="https://ajax.googleapis.com/ajax/libs/jqueryui/1.10.0/themes/cupertino/jquery-ui.css" type="text/css" media="all">

					  <link rel="icon" type="image/x-icon" href="https://assets.nationbuilder.com/springtide/sites/38/favicon_images/original/faviconcc.png?1496167237" />


				  <script src="https://assets.nationbuilder.com/assets/liquid/main-f52182358767f5af49bb34ddeeedb502a15f0105ffc14c758599b2789870803b.js"></script>

				<script type="text/javascript">
				  window.twttr = (function (d,s,id) {
					var t, js, fjs = d.getElementsByTagName(s)[0];
					if (d.getElementById(id)) return; js=d.createElement(s); js.id=id;
					js.src="//platform.twitter.com/widgets.js"; fjs.parentNode.insertBefore(js, fjs);
					return window.twttr || (t = { _e: [], ready: function(f){ t._e.push(f) } });
				  }(document, "script", "twitter-wjs"));
				</script>

				<script type="text/javascript">
				  NB.FBAppId = '1679673749184434';
				</script>

				  <script type="text/javascript">
					(function (d) { var config = { kitId: 'atg5ome', scriptTimeout: 3000, async: true }, h=d.documentElement,t=setTimeout(function(){h.className=h.className.replace(/\bwf-loading\b/g,"")+" wf-inactive";},config.scriptTimeout),tk=d.createElement("script"),f=false,s=d.getElementsByTagName("script")[0],a;h.className+=" wf-loading";tk.src='https://use.typekit.net/'+config.kitId+'.js';tk.async=true;tk.onload=tk.onreadystatechange=function(){a=this.readyState;if(f||a&&a!="complete"&&a!="loaded")return;f=true;clearTimeout(t);try{Typekit.load(config)}catch(e){}};s.parentNode.insertBefore(tk,s)})(document);
				  </script>

				  <script type="text/javascript">
					var _gaq = _gaq || [];
					_gaq.push(['_setAccount', 'UA-100048182-1']);
					 _gaq.push(['_setDomainName', 'none']);
					_gaq.push(['_gat._anonymizeIp']);
					_gaq.push(['_setAllowLinker', true]);
					  _gaq.push(['_setCustomVar', 1, 'UGC', 'false', 3]);
					  _gaq.push(['_setCustomVar', 1, 'Page type', 'Donation (v2)', 3]);
					_gaq.push(['_trackPageview']);
					_gaq.push(['_trackPageLoadTime']);

					(function() {
					  var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
						  ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';

					  let s = document.getElementsByTagName('script')[0];
					  s.parentNode.insertBefore(ga, s);
					})();
				  </script>






				  <div id='recaptcha-input'>
					<script src="https://www.recaptcha.net/recaptcha/api.js?render=6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T"   ></script>
						<script>
						  // Define function so that we can call it again later if we need to reset it
						  // This executes reCAPTCHA and then calls our callback.
						  function executeRecaptchaForDonate() {
							grecaptcha.ready(function() {
							  grecaptcha.execute('6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T', {action: 'donate'}).then(function(token) {
								setInputWithRecaptchaResponseTokenForDonate('g-recaptcha-response-data-donate', token)
							  });
							});
						  };
						  // Invoke immediately
						  executeRecaptchaForDonate()

						  // Async variant so you can await this function from another async function (no need for
						  // an explicit callback function then!)
						  // Returns a Promise that resolves with the response token.
						  async function executeRecaptchaForDonateAsync() {
							return new Promise((resolve, reject) => {
							  grecaptcha.ready(async function() {
								resolve(await grecaptcha.execute('6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T', {action: 'donate'}))
							  });
							})
						  };

									var setInputWithRecaptchaResponseTokenForDonate = function(id, token) {
							var element = document.getElementById(id);
							element.value = token;
						  }

						</script>
				<input type="hidden" name="g-recaptcha-response-data[donate]" id="g-recaptcha-response-data-donate" data-sitekey="6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T" class="g-recaptcha g-recaptcha-response "/>

					<script>
					  const clientKey = '6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T'
					  grecaptcha.ready( function() {
						var $badge = $('.grecaptcha-badge');
						if ($badge.length === 1){
						  if ($('#cd-nav, #control-panel-nav').length === 1) {
							$badge.css("bottom", "84px")
						  }
						  $badge.css("z-index", "9994")
						  $badge.find('iframe').css("margin", "0");
						}

						setInterval(
						  function() {
							grecaptcha.execute(clientKey, { action: 'donate' }).then(
							  function (token) {
								if ($("#g-recaptcha-response-data-donate").length) {
								  $("#g-recaptcha-response-data-donate").val(token);
								}
							  }
							)
						  }, 90000)
					  });
					</script>
				</div>


				  <script src="https://js.stripe.com/v3/"></script>
				  <script>
					NB.payments.publishableKey = "pk_live_1TjMHnI0fp51k3hKVhJEpm6D";
					NB.payments.epoButtonTheme = "dark";



					NB.payments.descriptor = "SPRINGTIDE"
				  </script>
					<script src="https://assets.nationbuilder.com/assets/payments_styling-3cb98755e69cd4399888137a7648ee6eabdd63c4d4c547626e7174d769be6e1b.js"></script>
						<script>
						  function initializeDefaultElementOptionsForCustomNationSignupPages() {}
						</script>
				  <script src="https://assets.nationbuilder.com/assets/payments-e2d51187f660cba449c70f7f87f733f762042cf340a5fd57e0f8c1fbf885c033.js"></script>

				<!-- The following line of CSS hides a checkbox in social share prompts for posting to Facebook;
				As of Aug 1 2018, FB has deprecated the ability for apps to post to personal profile pages, so this is a temporary fix
				while we re-configure the social share prompt -->
				<style>label[for="face_tweet_is_facebook"]{display:none;}</style>



				  <script type="text/javascript">
					NB.appConfig.userIsLoggedIn = false;
				  </script>

				  <link href='//fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic,700italic|Oswald:400,300,700|Lato:400,700,400italic,700italic|Bree+Serif|Josefin+Sans:700|Bitter:400,700,400italic|Open+Sans:400italic,700italic,400,700,800' rel='stylesheet' type='text/css'>
				<link href="https://fonts.googleapis.com/css?family=Open+Sans:400,400i,700,700i" rel="stylesheet">
				<link href="https://fonts.googleapis.com/css?family=Lato:900" rel="stylesheet">
				<!-- Facebook Pixel Code -->
				<script>
				  !function(f,b,e,v,n,t,s)
				  {if(f.fbq)return;n=f.fbq=function(){n.callMethod?
				  n.callMethod.apply(n,arguments):n.queue.push(arguments)};
				  if(!f._fbq)f._fbq=n;n.push=n;n.loaded=!0;n.version='2.0';
				  n.queue=[];t=b.createElement(e);t.async=!0;
				  t.src=v;s=b.getElementsByTagName(e)[0];
				  s.parentNode.insertBefore(t,s)}(window, document,'script',
				  'https://connect.facebook.net/en_US/fbevents.js');
				  fbq('init', '119846658737416');
				  fbq('track', 'PageView');
				</script>
				<noscript><img height="1" width="1" style="display:none" src="https://www.facebook.com/tr?id=119846658737416&ev=PageView&noscript=1" /></noscript>
				<!-- End Facebook Pixel Code -->

				<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/jquery.ui.effect.min.js"></script>
				<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/jquery.ui.effect-slide.min.js"></script>
				<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431716941215/default/staged-donations.js"></script>

				</head>
				<body class="aware-theme v2-theme page-type-donation-v2 js with-background">
				 <!-- Google Tag Manager (noscript) -->
				<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-P6TB763" height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
				<!-- End Google Tag Manager (noscript) -->


				  <div id="pattern" class="pattern">
					<div class="wrap" id="wrap">
					  <div id="body" class="page-pages-show-donation-v2-wide">

						<div class="header-container clearfix">
						  <div class="width-container clearfix">

							<div class="tablet-visible">

				<!-- _nav.html -->


				<!-- /_nav.html -->

							</div>


							  <div class="site-logo">
								<header><a href="/"><img src="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891"></a></header>
							  </div>



						  </div>
						  <!-- .width-container -->
						</div>
						<!-- .header-container -->

						<div class="nav-container desktop-visible">
						  <div class="width-container clearfix">

				<!-- _nav.html -->


				<!-- /_nav.html -->

						  </div>
						</div>




						<div class="main-container" id="middle">
						  <div class="main width-container clearfix">



				<!-- _columns_1.html -->
				<div class="onecolumn-container clearfix">

				  <div class="columns-1-flash"><div id="flash_container">



				</div>
				</div>

				  <div class="content-pages-show-donation-v2-wide">
					<div id="content">
				  <form id="donate_appeal_page_new_donation_form" class="donation_form" method="POST" action="/forms/donations"><input name="authenticity_token" type="hidden" value="vRw8PxR5c7Xb0QsfmIB9TVLL3SWNVx/ZK2sugu3oukI="/><input name="page_id" type="hidden" value="1780"/><input name="return_to" type="hidden" value="https://www.charterchallenge.ca/donate_appeal"/><div class="email_address_form" style="display:none;" aria-hidden="true"><p><label for="email_address">Optional email code</label><br/><input name="email_address" type="text" class="text" id="email_address" autocomplete="off"/></p></div>
				  <div class="form-wrap">

					  <div id="headline">
						<h2>Appeal Fund Stage 2 - Going to Court Again!</h2>
					  </div>


					<div id="intro" class="intro">
					  <p>Thanks to our very generous donors, we raised the full $25,000 needed for us to file our Notice of Appeal and to prepare and submit our written argument (<em>factum</em>). This also enabled us to claim the full $10,000 matching donation that a generous donor offered us.</p>
				<p>Over the past four years, over 1200 Canadians have collectively contributed almost $250,000 to get us this far - all the way from filing our Notice of Application in 2019 through to appearing in the Ontario Superior Court in September 2023.</p>
				<p>We're now inviting contributions towards the remaining $25,000 we estimate will be needed to prepare for and appear in the Appeal Court for Ontario - currently scheduled to happen in November 2024. These funds will be used to cover case management meetings with the judge, working with intervenors (organizations who will submit briefs in support of our challenge, such as Fair Vote Canada and Apathy is Boring), preparing for oral arguments, and actually appearing in court.</p>
				<p><span style="font-weight: 400;">To support our appeal, you can make a secure online donation via credit card below. To donate by cheque or e-transfer, please visit <a href="https://www.charterchallenge.ca/ways-to-give">this page</a>.</span></p>
					</div>






					  <div class="clearfix">

						<div class="progress" style="width: 100%;">
						  <div class="bar bar-success" style="width: 74.89248%;">

								<div class="bar-text">$18,723.12 raised*</div>

						  </div>
						</div>


						<div class="bar-goal">GOAL: $25,000.00</div>

					  </div>



					<div class="form">

					  <div class="form-errors">

					  </div>



					  <div class="row-fluid">
						<div class="span12">

						  <div class="row-fluid padbottomless">
							<h4>1. Amount</h4>
						  </div>
						  <div class="row-fluid">
							<div class="span12">

								<div class="radio-inline donation-v2-amounts padbottommore">
				  <span>

					  <input id="donation_amount_25"
							 type="radio"
							 name="donation[amount_option]"
							 class="donation_amount_option"
							 value="25"
							  />
					<label for="donation_amount_25" class="radio">$25</label>
				  </span>

				  <span>

					  <input id="donation_amount_50"
							 type="radio"
							 name="donation[amount_option]"
							 class="donation_amount_option"
							 value="50"
							  />
					<label for="donation_amount_50" class="radio">$50</label>
				  </span>

				  <span>

					  <input id="donation_amount_100"
							 type="radio"
							 name="donation[amount_option]"
							 class="donation_amount_option"
							 value="100"
							  />
					<label for="donation_amount_100" class="radio">$100</label>
				  </span>

				  <span>

					  <input id="donation_amount_250"
							 type="radio"
							 name="donation[amount_option]"
							 class="donation_amount_option"
							 value="250"
							  />
					<label for="donation_amount_250" class="radio">$250</label>
				  </span>

				  <span>

					  <input id="donation_amount_1000"
							 type="radio"
							 name="donation[amount_option]"
							 class="donation_amount_option"
							 value="1000"
							  />
					<label for="donation_amount_1000" class="radio">$1,000</label>
				  </span>

				  <span>

					  <input id="donation_amount_2500"
							 type="radio"
							 name="donation[amount_option]"
							 class="donation_amount_option"
							 value="2500"
							  />
					<label for="donation_amount_2500" class="radio">$2,500</label>
				  </span>

				<span>
				  <input id="donation_amount_other"
						 type="radio"
						 name="donation[amount_option]"
						 class="donation_amount_option"
						  />
				  <label for="donation_amount_other" class="radio">Other</label>
				</span>
				</div>



							  <div class="row-fluid donation-v2-options ">

								<div class="span6">
								  <div class="donation-other-input-container">
				  <div class="currency-symbol">$</div>
				  <input id="donation_amount_other_input"
						 type="text"
						 name="donation[amount]"
						 class="text nb_donation_v2_amount nb_donation_v2_amount_$"
						 />
				</div>

								</div>


								<div class="span6">
								  <div class="donation-v2-occurence-radio">
									<span>
				  <input id="donation_donation_occurrence_one_time"
						 type="radio"
						 name="donation[donation_occurrence]"
						 class="donation_amount_option"
						 value="one-time"
						 checked="checked" />
				  <label class="radio" for="donation_donation_occurrence_one_time">
					 One-time
				  </label>
				</span>
				<span>
				  <input id="donation_donation_occurrence_monthly"
						type="radio"
						name="donation[donation_occurrence]"
						class="donation_amount_option"
						value="monthly"
						 />
				  <label class="radio" for="donation_donation_occurrence_monthly">
					Monthly
				  </label>
				</span>

								  </div>
								</div>

							  </div>




				When you choose an amount, please remember that all donations above $25 will generate a charitable donation tax receipt. The value of these credits varies by province or territory, ranging from a 19-35% refund for the first $200 donated to a 40-53% refund on any amounts donated above $200.

							</div>
						  </div>

						  <div class="row-fluid padbottomless">
							<h4 class="sub-header">2. Your information</h4>
						  </div>
						  <div class="row-fluid">
							<div class="span12"><label for="donation_email">Email</label><input class="text" id="donation_email" name="donation[email]" type="email" /></div>
						  </div>
						  <div class="row-fluid">
							<div class="span6"><label for="donation_first_name">First Name</label><input class="text" id="donation_first_name" name="donation[first_name]" type="text" /></div>
							<div class="span6"><label for="donation_last_name">Last Name</label><input class="text" id="donation_last_name" name="donation[last_name]" type="text" /></div>

						  </div>
						  <div class="row-fluid">

							<div class="span12">
							  <label for="donation_billing_address_country_code">Country</label><select id="donation_billing_address_country_code" name="donation[billing_address_attributes][country_code]" class="select"><option value="AF">Afghanistan</option><option value="AL">Albania</option><option value="DZ">Algeria</option><option value="AS">American Samoa</option><option value="AD">Andorra</option><option value="AO">Angola</option><option value="AI">Anguilla</option><option value="AQ">Antarctica</option><option value="AG">Antigua and Barbuda</option><option value="AR">Argentina</option><option value="AM">Armenia</option><option value="AW">Aruba</option><option value="AU">Australia</option><option value="AT">Austria</option><option value="AZ">Azerbaijan</option><option value="BS">Bahamas</option><option value="BH">Bahrain</option><option value="BD">Bangladesh</option><option value="BB">Barbados</option><option value="BY">Belarus</option><option value="BE">Belgium</option><option value="BZ">Belize</option><option value="BJ">Benin</option><option value="BM">Bermuda</option><option value="BT">Bhutan</option><option value="BO">Bolivia</option><option value="BQ">Bonaire, Sint Eustatius and Saba</option><option value="BA">Bosnia and Herzegovina</option><option value="BW">Botswana</option><option value="BV">Bouvet Island</option><option value="BR">Brazil</option><option value="IO">British Indian Ocean Territory</option><option value="BN">Brunei Darussalam</option><option value="BG">Bulgaria</option><option value="BF">Burkina Faso</option><option value="BI">Burundi</option><option value="KH">Cambodia</option><option value="CM">Cameroon</option><option value="CA" selected="selected">Canada</option><option value="CV">Cape Verde</option><option value="KY">Cayman Islands</option><option value="CF">Central African Republic</option><option value="TD">Chad</option><option value="CL">Chile</option><option value="CN">China</option><option value="CX">Christmas Island</option><option value="CC">Cocos (Keeling) Islands</option><option value="CO">Colombia</option><option value="KM">Comoros</option><option value="CG">Congo</option><option value="CD">Congo, the Democratic Republic of the</option><option value="CK">Cook Islands</option><option value="CR">Costa Rica</option><option value="HR">Croatia</option><option value="CU">Cuba</option><option value="CW">Curaçao</option><option value="CY">Cyprus</option><option value="CZ">Czech Republic</option><option value="CI">Côte d'Ivoire</option><option value="DK">Denmark</option><option value="DJ">Djibouti</option><option value="DM">Dominica</option><option value="DO">Dominican Republic</option><option value="EC">Ecuador</option><option value="EG">Egypt</option><option value="SV">El Salvador</option><option value="GQ">Equatorial Guinea</option><option value="ER">Eritrea</option><option value="EE">Estonia</option><option value="ET">Ethiopia</option><option value="FK">Falkland Islands (Malvinas)</option><option value="FO">Faroe Islands</option><option value="FJ">Fiji</option><option value="FI">Finland</option><option value="FR">France</option><option value="GF">French Guiana</option><option value="PF">French Polynesia</option><option value="TF">French Southern Territories</option><option value="GA">Gabon</option><option value="GM">Gambia</option><option value="GE">Georgia</option><option value="DE">Germany</option><option value="GH">Ghana</option><option value="GI">Gibraltar</option><option value="GR">Greece</option><option value="GL">Greenland</option><option value="GD">Grenada</option><option value="GP">Guadeloupe</option><option value="GU">Guam</option><option value="GT">Guatemala</option><option value="GG">Guernsey</option><option value="GN">Guinea</option><option value="GW">Guinea-Bissau</option><option value="GY">Guyana</option><option value="HT">Haiti</option><option value="HM">Heard Island and McDonald Islands</option><option value="VA">Holy See (Vatican City State)</option><option value="HN">Honduras</option><option value="HK">Hong Kong</option><option value="HU">Hungary</option><option value="IS">Iceland</option><option value="IN">India</option><option value="ID">Indonesia</option><option value="IR">Iran, Islamic Republic of</option><option value="IQ">Iraq</option><option value="IE">Ireland</option><option value="IM">Isle of Man</option><option value="IL">Israel</option><option value="IT">Italy</option><option value="JM">Jamaica</option><option value="JP">Japan</option><option value="JE">Jersey</option><option value="JO">Jordan</option><option value="KZ">Kazakhstan</option><option value="KE">Kenya</option><option value="KI">Kiribati</option><option value="KW">Kuwait</option><option value="KG">Kyrgyzstan</option><option value="LA">Lao People's Democratic Republic</option><option value="LV">Latvia</option><option value="LB">Lebanon</option><option value="LS">Lesotho</option><option value="LR">Liberia</option><option value="LY">Libya</option><option value="LI">Liechtenstein</option><option value="LT">Lithuania</option><option value="LU">Luxembourg</option><option value="MO">Macao</option><option value="MG">Madagascar</option><option value="MW">Malawi</option><option value="MY">Malaysia</option><option value="MV">Maldives</option><option value="ML">Mali</option><option value="MT">Malta</option><option value="MH">Marshall Islands</option><option value="MQ">Martinique</option><option value="MR">Mauritania</option><option value="MU">Mauritius</option><option value="YT">Mayotte</option><option value="MX">Mexico</option><option value="FM">Micronesia, Federated States of</option><option value="MD">Moldova, Republic of</option><option value="MC">Monaco</option><option value="MN">Mongolia</option><option value="ME">Montenegro</option><option value="MS">Montserrat</option><option value="MA">Morocco</option><option value="MZ">Mozambique</option><option value="MM">Myanmar</option><option value="NA">Namibia</option><option value="NR">Nauru</option><option value="NP">Nepal</option><option value="NL">Netherlands</option><option value="NC">New Caledonia</option><option value="NZ">New Zealand</option><option value="NI">Nicaragua</option><option value="NE">Niger</option><option value="NG">Nigeria</option><option value="NU">Niue</option><option value="NF">Norfolk Island</option><option value="KP">North Korea</option><option value="MK">North Macedonia, Republic of</option><option value="MP">Northern Mariana Islands</option><option value="NO">Norway</option><option value="OM">Oman</option><option value="PK">Pakistan</option><option value="PW">Palau</option><option value="PS">Palestine, State of</option><option value="PA">Panama</option><option value="PG">Papua New Guinea</option><option value="PY">Paraguay</option><option value="PE">Peru</option><option value="PH">Philippines</option><option value="PN">Pitcairn</option><option value="PL">Poland</option><option value="PT">Portugal</option><option value="PR">Puerto Rico</option><option value="QA">Qatar</option><option value="RO">Romania</option><option value="RU">Russian Federation</option><option value="RW">Rwanda</option><option value="RE">Réunion</option><option value="BL">Saint Barthélemy</option><option value="SH">Saint Helena, Ascension and Tristan da Cunha</option><option value="KN">Saint Kitts and Nevis</option><option value="LC">Saint Lucia</option><option value="MF">Saint Martin (French part)</option><option value="PM">Saint Pierre and Miquelon</option><option value="VC">Saint Vincent and the Grenadines</option><option value="WS">Samoa</option><option value="SM">San Marino</option><option value="ST">Sao Tome and Principe</option><option value="SA">Saudi Arabia</option><option value="SN">Senegal</option><option value="RS">Serbia</option><option value="SC">Seychelles</option><option value="SL">Sierra Leone</option><option value="SG">Singapore</option><option value="SX">Sint Maarten (Dutch part)</option><option value="SK">Slovakia</option><option value="SI">Slovenia</option><option value="SB">Solomon Islands</option><option value="SO">Somalia</option><option value="ZA">South Africa</option><option value="GS">South Georgia and the South Sandwich Islands</option><option value="KR">South Korea</option><option value="SS">South Sudan</option><option value="ES">Spain</option><option value="LK">Sri Lanka</option><option value="SD">Sudan</option><option value="SR">Suriname</option><option value="SJ">Svalbard and Jan Mayen</option><option value="SZ">Swaziland</option><option value="SE">Sweden</option><option value="CH">Switzerland</option><option value="SY">Syrian Arab Republic</option><option value="TW">Taiwan</option><option value="TJ">Tajikistan</option><option value="TZ">Tanzania, United Republic of</option><option value="TH">Thailand</option><option value="TL">Timor-Leste</option><option value="TG">Togo</option><option value="TK">Tokelau</option><option value="TO">Tonga</option><option value="TT">Trinidad and Tobago</option><option value="TN">Tunisia</option><option value="TR">Turkey</option><option value="TM">Turkmenistan</option><option value="TC">Turks and Caicos Islands</option><option value="TV">Tuvalu</option><option value="UG">Uganda</option><option value="UA">Ukraine</option><option value="AE">United Arab Emirates</option><option value="GB">United Kingdom</option><option value="US">United States</option><option value="UM">United States Minor Outlying Islands</option><option value="UY">Uruguay</option><option value="UZ">Uzbekistan</option><option value="VU">Vanuatu</option><option value="VE">Venezuela, Bolivarian Republic of</option><option value="VN">Viet Nam</option><option value="VG">Virgin Islands, British</option><option value="VI">Virgin Islands, U.S.</option><option value="WF">Wallis and Futuna</option><option value="EH">Western Sahara</option><option value="YE">Yemen</option><option value="ZM">Zambia</option><option value="ZW">Zimbabwe</option><option value="AX">Åland Islands</option></select>
							</div>

						  </div>
						  <div class="row-fluid">

							<div class="span12"><label for="donation_billing_address_address1">Address</label>
							  <input class="text" id="donation_billing_address_address1" name="donation[billing_address_attributes][address1]" type="text" />
							  <input class="text" id="donation_billing_address_address2" name="donation[billing_address_attributes][address2]" type="text" />
							  <input class="text not-us-or-canada hide" id="donation_billing_address_address3" name="donation[billing_address_attributes][address3]" type="text" />
							</div>

						  </div>
						  <div class="row-fluid">
							<div class="span4">
							  <label for="donation_billing_address_city">City/Town</label><input class="text" id="donation_billing_address_city" name="donation[billing_address_attributes][city]" type="text" />
							</div>
							<div class="span4 us-or-canada us-only hide">
							  <label for="donation_state">Province</label><select id="donation_billing_address_state" name="donation[billing_address_attributes][state]" class="select"><option value="" selected="selected"></option><option value="AL">Alabama</option><option value="AK">Alaska</option><option value="AS">American Samoa</option><option value="AZ">Arizona</option><option value="AR">Arkansas</option><option value="AA">Armed Forces Americas</option><option value="AE">Armed Forces Europe</option><option value="AP">Armed Forces Pacific</option><option value="CA">California</option><option value="CO">Colorado</option><option value="CT">Connecticut</option><option value="DE">Delaware</option><option value="DC">District of Columbia</option><option value="FM">Federated States of Micronesia</option><option value="FL">Florida</option><option value="GA">Georgia</option><option value="GU">Guam</option><option value="HI">Hawaii</option><option value="ID">Idaho</option><option value="IL">Illinois</option><option value="IN">Indiana</option><option value="IA">Iowa</option><option value="KS">Kansas</option><option value="KY">Kentucky</option><option value="LA">Louisiana</option><option value="ME">Maine</option><option value="MH">Marshall Islands</option><option value="MD">Maryland</option><option value="MA">Massachusetts</option><option value="MI">Michigan</option><option value="MN">Minnesota</option><option value="MS">Mississippi</option><option value="MO">Missouri</option><option value="MT">Montana</option><option value="NE">Nebraska</option><option value="NV">Nevada</option><option value="NH">New Hampshire</option><option value="NJ">New Jersey</option><option value="NM">New Mexico</option><option value="NY">New York</option><option value="NC">North Carolina</option><option value="ND">North Dakota</option><option value="OH">Ohio</option><option value="OK">Oklahoma</option><option value="OR">Oregon</option><option value="PW">Palau</option><option value="PA">Pennsylvania</option><option value="PR">Puerto Rico</option><option value="RI">Rhode Island</option><option value="SC">South Carolina</option><option value="SD">South Dakota</option><option value="TN">Tennessee</option><option value="TX">Texas</option><option value="UT">Utah</option><option value="VT">Vermont</option><option value="VI">Virgin Island</option><option value="VA">Virginia</option><option value="WA">Washington</option><option value="WV">West Virginia</option><option value="WI">Wisconsin</option><option value="WY">Wyoming</option></select>
							</div>
							<div class="span4 us-or-canada canada-only hide">
							  <label for="donation_billing_address_state">Province</label><input class="text" id="donation_billing_address_state" name="donation[billing_address_attributes][state]" type="text" />
							</div>
							<div class="span4">
							  <label for="donation_billing_address_zip">Postal Code</label><input class="text" id="donation_billing_address_zip" name="donation[billing_address_attributes][zip]" type="text" />
							</div>

						  </div>
						  <div class="row-fluid">

							<div class="span12">
							  <label for="donation_billing_address_phone_number">Phone Number</label><input class="text" id="donation_billing_address_phone_number" name="donation[billing_address_attributes][phone_number]" type="tel" />
							</div>

						  </div>

						  <div class="row-fluid">

							<div class="span12">
							  <label class="checkbox" for="donation_email_opt_in"><input name="donation[email_opt_in]" type="hidden" value="0" /><input checked="checked" id="donation_email_opt_in" name="donation[email_opt_in]" type="checkbox" value="1" />
				Send me email updates from Springtide, and allow both Springtide and Fair Voting BC to contact me for the purposes of communication about the Charter Challenge for Fair Voting Project.
				<br>
								<br>
								(If you have already opted into emails from Springtide and/or the Charter Challenge project, unchecking this box will mean you will no longer receive any updates from us).
							  </label>
							</div>

						  </div>








						  <div class="row-fluid padbottomless">
							<h4 class="sub-header">3. Payment information</h4>
						  </div>
						  <div class="row-fluid">
							<div class="span12">
							  <label>Credit Card</label>
							  <div class="payment-input payment-toggle-view" data-payments-element-type="card" data-error-container=".payment-input + div.card-error-container"></div>
				<div class="card-error-container"></div>

							</div>
						  </div>
						  <div class="row-fluid padtopless padbottomless padtop">
							<div clas="span12">


							  <label for="donation_is_private" class="checkbox padtopmore"><input name="donation[is_private]" type="hidden" value="0" /><input class="checkbox" id="donation_is_private" name="donation[is_private]" type="checkbox" value="1" /> Don't publish my donation on the website.</label>

							  <span style="font-size: xx-small;">

							  Note: all donations processed on this page are secure payments processed via <a href="https://stripe.com/en-ca">Stripe</a>, which has the highest level of security certification attainable in the payments industry (PCI Level 1). For alternative ways to donate, visit <a href="https://www.charterchallenge.ca/ways-to-give"> this page</a>.
							  </span>
							  </div>
						  </div>
						  <div class="row-fluid">
							<div class="span12">

							</div>
							<div class="span12">
							  <div class="submit-container">
								<div class="donation-v2-amount">

								  <span>

									<span class="hidden">$</span><span class="nb_donation_v2_amount">Please select an amount</span>


									<div class="nb_donation_v2_interval" data-placeholder="paid monthly"></div>

								  </span>

								</div>

								<input class="submit-button btn btn-primary btn-lg" type="submit" name="commit" value="Donate now" />

							  </div>
							</div>
							<div class="form-submit"></div>
						  </div>
						</div>

					  </div>

					</div>
				  </div>

				  </form>

				</div>

				<span style="font-size: xx-small;">
				  <p> * Total reflects processed donations only. Donations made from this page are processed immediately, however some donations (e.g. cheque, e-transfer, paypal) are processed manually and there can be a delay (of up to 1 - 2 weeks) before they will be added to this total.</p>
				  <p> All donations to the Charter Challenge for Fair Voting are made to the Springtide Collective for Democracy Society -  P.O. BOX 40003| Robie Street PO | Halifax, NS Canada | B3K 0E4 | 902.989.3668
					<p>Springtide is a registered Canadian charity. Our CRA charitable registration number is 838267136RR0001. For more on charitable giving in Canada, see <a href=https://www.canada.ca/charities-giving>canada.ca/charities-giving</a>
				<p> <br>Contact us: <a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="630a0d050c231013110a0d04170a07064d0d040c">[email&#160;protected]</a>
				  <p>In the event that the fundraising goal is not met, Springtide (the recipient charity of all donations) retains the right to hold the funds collected for use in a future case, or alternatively, to create educational programming and resources for teaching and learning about electoral reform in Canada, or other charitable purposes as Springtide sees fit.
				</span>

				<script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script src="/assets/liquid/theme_donation_v2.js"></script>

				  </div>

				</div>
				<!-- /onecolumn-container -->
				<!-- /_columns_1.html -->



						  </div>
						  <!-- .main -->
						</div>
						<!-- .main-container -->

						<footer>
						  <div class="width-container clearfix">

							<div id="fb-root"></div>

				<script type="text/javascript">
				  window.fbAsyncInit = function() {
					FB.init({
					  appId      : 1679673749184434,
					  channelUrl : "//cc-springtide.nationbuilder.com/channel.html",
					  status     : true,
					  version    : "v18.0",
					  cookie     : true,
					  xfbml      : true
					});
				  };
				  (function(d, s, id) {
					var js, fjs = d.getElementsByTagName(s)[0];
					if (d.getElementById(id)) return;
					js = d.createElement(s);
					js.id = id;
					js.async = true;
					js.src = "//connect.facebook.net/en_US/sdk.js";
					fjs.parentNode.insertBefore(js, fjs);
				  }(document, 'script', 'facebook-jssdk'));
				</script>

				<script src="https://assets.nationbuilder.com/assets/liquid-6fda76e47cd1a46bec92e2adac0a0453c78638197e234d7667c2ff4366c5a44a.js"></script>








							<div class="row-fluid">

							  <div class="span6">

							  </div>



							</div>

						  </div>
						  <!-- .width-container -->
						</footer>

					  </div>
					  <!-- #body -->
					</div>
					<!-- #wrap -->
				  </div>
				  <!-- #pattern -->




				  <!--[if lt IE 9]>
				  <script type="text/javascript" src="/javascripts/jquery.backstretch.min.js"></script>
				  <script type="text/javascript">
					jQuery.backstretch("https://assets.nationbuilder.com/springtide/pages/1780/header_images/original/gavel.png?1701407475", {speed: 0});
				  </script>
				  <![endif]-->
				  <style>
					body {
					  background: url('https://assets.nationbuilder.com/springtide/pages/1780/header_images/original/gavel.png?1701407475') no-repeat center center fixed;
					  -webkit-background-size: cover;
					  -moz-background-size: cover;
					  -o-background-size: cover;
					  background-size: cover;
					}
				</style>


				<script type="text/javascript">
				  function pollFBCommentBox(widget) {
					var fbCommentBoxTimer = window.setInterval(function() {
					  if($(widget).height() < 100) {
						$(widget).removeClass('comment-box-active');
						window.clearInterval(fbCommentBoxTimer);
					  }
					}, 1000);
				  }
				  FB.Event.subscribe('edge.create', function(href, widget) {
					$(widget).addClass('comment-box-active');
					pollFBCommentBox(widget);
				  });
				  FB.Event.subscribe('edge.remove', function(href, widget) {
					$(widget).removeClass('comment-box-active');
				  });
				</script>

				  <script>
					if ( window.self !== window.top ) {
					  var referrer_origin = "";
					  if ( window.location.origin !== referrer_origin ) {
						var xhttp = new XMLHttpRequest();
						var params = "iframe_req_path=" + window.location.pathname + "&referrer_origin=" + referrer_origin;
						xhttp.open("GET", "iframe_requests?" + params, true);
						xhttp.send();
					  }
					}
				  </script>
				<script defer src="https://static.cloudflareinsights.com/beacon.min.js/vef91dfe02fce4ee0ad053f6de4f175db1715022073587" integrity="sha512-sDIX0kl85v1Cl5tu4WGLZCpH/dV9OHbA4YlKCuCiMmOQIk4buzoYDZSFj+TvC71mOBLh8CDC/REgE0GX0xcbjA==" data-cf-beacon='{"rayId":"88db4b652e8baac5","version":"2024.4.1","token":"ff633690ee944cc4bb4fa2453466614e"}' crossorigin="anonymous"></script>
				</body>

				</html>


				`,
			},
			want:    "$25,000.00",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseGoal(tt.args.pageContents)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGoal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseGoal() = %v, want %v", got, tt.want)
			}
		})
	}
}

const webPageWith5526 = `
<!DOCTYPE html>
<!--[if lt IE 7]>
<html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
<!--[if IE 7]>
<html class="no-js lt-ie9 lt-ie8"> <![endif]-->
<!--[if IE 8]>
<html class="no-js lt-ie9"> <![endif]-->
<!--[if gt IE 8]><!-->
<html class="no-js"> <!--<![endif]-->
<head>
  <meta name="google-site-verification" content="pJgtWPNLX5_ucuNccfFyJu8LSQ-aDzrnVveSiUc23x8" />
  <meta name="google-site-verification" content="CjNDrfIadB9PJb8DiRI7IDIC2Ud_s4OdC1A9eLfNXP8" />
  <meta name="google-site-verification" content="tWVq_LM8h4MFZ51wf5z6L8SzXyA115_g_ie4MFOKijg" />
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <title>Donate - Charter Challenge for Fair Voting</title>
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css" rel="stylesheet">
  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431671722882/default/theme.scss" type="text/css" />
  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431671722882/default/tablet-and-desktop.scss" type="text/css" media="screen and (min-width: 768px)" />

  <!-- because ie8 ignores media queries, we need this -->
  <!--[if IE 8]>
    <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431671722882/default/tablet-and-desktop.scss" type="text/css" />
  <![endif]-->


  <!--[if IE]>
  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431671722882/default/ie.scss" type="text/css" />
  <![endif]-->



<script type="text/javascript">var _sf_startpt=(new Date()).getTime()</script>

<meta content="authenticity_token" name="csrf-param" />
<meta content="XQW22+kxiOQUSKfmXngiN1hV5HUTRmrDoIP/UkMuND0=" name="csrf-token" />

  <link rel="canonical" href="https://www.charterchallenge.ca/donate_charter_challenge" />
    <meta name="Title" content="Donate">
    <meta property="og:title" content="Donate"/>
  <meta property="og:url" content="https://www.charterchallenge.ca/donate_charter_challenge">
    <meta property="og:type" content="article">
      <link rel="image_src" href="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891" />
      <meta property="og:image" content="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891" />
  <meta property="og:site_name" content="Charter Challenge for Fair Voting"/>

<script type="text/javascript">
  var NB = NB || {};

  NB.environment = "production";
  NB.pageId = "1773";
  NB.Liquid = NB.Liquid || {
    Theme: {
      version: 2,
      name: "Charter Challenge for Fair Voting",
      parent: "publish",
      variation: ""
    }
  };
  NB.payments = NB.payments || {elements: {}};
  NB.payments.decimal_mark = "."
  NB.payments.thousands_separator = ","
  NB.payments.currency = "cad";
</script>

<script type="text/javascript">
    //<![CDATA[
      window._auth_token_name = "authenticity_token";
      window._auth_token = "XQW22+kxiOQUSKfmXngiN1hV5HUTRmrDoIP/UkMuND0=";
    //]]>
</script>

  <link rel="stylesheet" href="https://ajax.googleapis.com/ajax/libs/jqueryui/1.10.0/themes/cupertino/jquery-ui.css" type="text/css" media="all">

      <link rel="icon" type="image/x-icon" href="https://assets.nationbuilder.com/springtide/sites/38/favicon_images/original/faviconcc.png?1496167237" />


  <script src="https://assets.nationbuilder.com/assets/liquid/main-ddd08e9e6a89697bf95bed251cd7280f3a9ea8447407c06ffc458884c536d760.js"></script>

<script type="text/javascript">
  window.twttr = (function (d,s,id) {
    var t, js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return; js=d.createElement(s); js.id=id;
    js.src="//platform.twitter.com/widgets.js"; fjs.parentNode.insertBefore(js, fjs);
    return window.twttr || (t = { _e: [], ready: function(f){ t._e.push(f) } });
  }(document, "script", "twitter-wjs"));
</script>

<script type="text/javascript">
  NB.FBAppId = '126739610711965';
</script>

  <script type="text/javascript">
    (function (d) { var config = { kitId: 'atg5ome', scriptTimeout: 3000, async: true }, h=d.documentElement,t=setTimeout(function(){h.className=h.className.replace(/\bwf-loading\b/g,"")+" wf-inactive";},config.scriptTimeout),tk=d.createElement("script"),f=false,s=d.getElementsByTagName("script")[0],a;h.className+=" wf-loading";tk.src='https://use.typekit.net/'+config.kitId+'.js';tk.async=true;tk.onload=tk.onreadystatechange=function(){a=this.readyState;if(f||a&&a!="complete"&&a!="loaded")return;f=true;clearTimeout(t);try{Typekit.load(config)}catch(e){}};s.parentNode.insertBefore(tk,s)})(document);
  </script>

  <script type="text/javascript">
    var _gaq = _gaq || [];
    _gaq.push(['_setAccount', 'UA-100048182-1']);
     _gaq.push(['_setDomainName', 'none']);
    _gaq.push(['_gat._anonymizeIp']);
    _gaq.push(['_setAllowLinker', true]);
      _gaq.push(['_setCustomVar', 1, 'UGC', 'false', 3]);
      _gaq.push(['_setCustomVar', 1, 'Page type', 'Donation (v2)', 3]);
    _gaq.push(['_trackPageview']);
    _gaq.push(['_trackPageLoadTime']);

    (function() {
      var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
          ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';

      let s = document.getElementsByTagName('script')[0];
      s.parentNode.insertBefore(ga, s);
    })();
  </script>






  <div id='recaptcha-input'>
    <script src="https://www.recaptcha.net/recaptcha/api.js?render=6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T"   ></script>
        <script>
          // Define function so that we can call it again later if we need to reset it
          // This executes reCAPTCHA and then calls our callback.
          function executeRecaptchaForDonate() {
            grecaptcha.ready(function() {
              grecaptcha.execute('6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T', {action: 'donate'}).then(function(token) {
                setInputWithRecaptchaResponseTokenForDonate('g-recaptcha-response-data-donate', token)
              });
            });
          };
          // Invoke immediately
          executeRecaptchaForDonate()

          // Async variant so you can await this function from another async function (no need for
          // an explicit callback function then!)
          // Returns a Promise that resolves with the response token.
          async function executeRecaptchaForDonateAsync() {
            return new Promise((resolve, reject) => {
              grecaptcha.ready(async function() {
                resolve(await grecaptcha.execute('6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T', {action: 'donate'}))
              });
            })
          };

                    var setInputWithRecaptchaResponseTokenForDonate = function(id, token) {
            var element = document.getElementById(id);
            element.value = token;
          }

        </script>
<input type="hidden" name="g-recaptcha-response-data[donate]" id="g-recaptcha-response-data-donate" data-sitekey="6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T" class="g-recaptcha g-recaptcha-response "/>

    <script>
      const clientKey = '6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T'
      grecaptcha.ready( function() {
        var $badge = $('.grecaptcha-badge');
        if ($badge.length === 1){
          if ($('#cd-nav, #control-panel-nav').length === 1) {
            $badge.css("bottom", "84px")
          }
          $badge.css("z-index", "9994")
          $badge.find('iframe').css("margin", "0");
        }

        setInterval(
          function() {
            grecaptcha.execute(clientKey, { action: 'donate' }).then(
              function (token) {
                if ($("#g-recaptcha-response-data-donate").length) {
                  $("#g-recaptcha-response-data-donate").val(token);
                }
              }
            )
          }, 90000)
      });
    </script>
</div>


  <script src="https://js.stripe.com/v3/"></script>
  <script>
    NB.payments.publishableKey = "pk_live_1TjMHnI0fp51k3hKVhJEpm6D";
    NB.payments.epoButtonTheme = "dark";



    NB.payments.descriptor = "SPRINGTIDE"
  </script>
    <script src="https://assets.nationbuilder.com/assets/payments_styling-ad9653f0ddd18fbd898a382fc934c3af9646949f19b2ba928321865cd6d5ebfb.js"></script>
        <script>
          function initializeDefaultElementOptionsForCustomNationSignupPages() {}
        </script>
  <script src="https://assets.nationbuilder.com/assets/payments-cde4ea419a4550e311d7fdf93d103a9f5f71d7f2789ce88d1f6cea589993e854.js"></script>
<!-- The following line of CSS hides a checkbox in social share prompts for posting to Facebook;
As of Aug 1 2018, FB has deprecated the ability for apps to post to personal profile pages, so this is a temporary fix
while we re-configure the social share prompt -->
<style>label[for="face_tweet_is_facebook"]{display:none;}</style>


  <script type="text/javascript">
    NB.appConfig.userIsLoggedIn = false;
  </script>

  <link href='//fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic,700italic|Oswald:400,300,700|Lato:400,700,400italic,700italic|Bree+Serif|Josefin+Sans:700|Bitter:400,700,400italic|Open+Sans:400italic,700italic,400,700,800' rel='stylesheet' type='text/css'>
<link href="https://fonts.googleapis.com/css?family=Open+Sans:400,400i,700,700i" rel="stylesheet">
<link href="https://fonts.googleapis.com/css?family=Lato:900" rel="stylesheet">
<!-- Facebook Pixel Code -->
<script>
  !function(f,b,e,v,n,t,s)
  {if(f.fbq)return;n=f.fbq=function(){n.callMethod?
  n.callMethod.apply(n,arguments):n.queue.push(arguments)};
  if(!f._fbq)f._fbq=n;n.push=n;n.loaded=!0;n.version='2.0';
  n.queue=[];t=b.createElement(e);t.async=!0;
  t.src=v;s=b.getElementsByTagName(e)[0];
  s.parentNode.insertBefore(t,s)}(window, document,'script',
  'https://connect.facebook.net/en_US/fbevents.js');
  fbq('init', '119846658737416');
  fbq('track', 'PageView');
</script>
<noscript><img height="1" width="1" style="display:none" src="https://www.facebook.com/tr?id=119846658737416&ev=PageView&noscript=1" /></noscript>
<!-- End Facebook Pixel Code -->

<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431671722882/default/jquery.ui.effect.min.js"></script>
<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431671722882/default/jquery.ui.effect-slide.min.js"></script>
<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431671722882/default/staged-donations.js"></script>

</head>
<body class="aware-theme v2-theme page-type-donation-v2 js with-background">
 <!-- Google Tag Manager (noscript) -->
<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-P6TB763" height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
<!-- End Google Tag Manager (noscript) -->


  <div id="pattern" class="pattern">
    <div class="wrap" id="wrap">
      <div id="body" class="page-pages-show-donation-v2-wide">

        <div class="header-container clearfix">
          <div class="width-container clearfix">

            <div class="tablet-visible">

<!-- _nav.html -->


<!-- /_nav.html -->

            </div>


              <div class="site-logo">
                <header><a href="/"><img src="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891"></a></header>
              </div>



          </div>
          <!-- .width-container -->
        </div>
        <!-- .header-container -->

        <div class="nav-container desktop-visible">
          <div class="width-container clearfix">

<!-- _nav.html -->


<!-- /_nav.html -->

          </div>
        </div>




        <div class="main-container" id="middle">
          <div class="main width-container clearfix">



<!-- _columns_1.html -->
<div class="onecolumn-container clearfix">

  <div class="columns-1-flash"><div id="flash_container">



</div>
</div>

  <div class="content-pages-show-donation-v2-wide">
    <div id="content">
  <form id="donate_charter_challenge_page_new_donation_form" class="donation_form" method="POST" action="/forms/donations"><input name="authenticity_token" type="hidden" value="XQW22+kxiOQUSKfmXngiN1hV5HUTRmrDoIP/UkMuND0="/><input name="page_id" type="hidden" value="1773"/><input name="return_to" type="hidden" value="https://www.charterchallenge.ca/donate_charter_challenge"/><div class="email_address_form" style="display:none;" aria-hidden="true"><p><label for="email_address">Optional email code</label><br/><input name="email_address" type="text" class="text" id="email_address" autocomplete="off"/></p></div>
  <div class="form-wrap">

      <div id="headline">
        <h2>Donate</h2>
      </div>


    <div id="intro" class="intro">
      <p>We need your support to fund the legal fees and supporting campaign for the Charter Challenge for Fair Voting - a court case against Canada's unfair voting system. </p>
<p>Over the past four years, we've raised over $180,000 that has brought us to the point where we have submitted our Notice of Application, our affidavit and evidence package, and starting preparing reply evidence. We are now seeking $25,000 to <span style="font-weight: 400;">cover the costs of presenting our case in court. </span></p>
<p><span>By supporting this case, you'll be joining the 1000+ Canadians who have already contributed enough support </span><span>to cover the costs of preparing a court filing and securing expert testimony for our case.</span></p>
<p><span style="font-weight: 400;">You can make a secure online donation via credit card, below. For alternative ways to donate to the Charter Challenge, visit <a href="https://www.charterchallenge.ca/ways-to-give">this page</a>.</span></p>
    </div>






      <div class="clearfix">

        <div class="progress" style="width: 100%;">
          <div class="bar bar-success" style="width: 22.104%;">

                <div class="bar-text">$5,526.00 raised*</div>

          </div>
        </div>


        <div class="bar-goal">GOAL: $25,000.00</div>

      </div>



    <div class="form">

      <div class="form-errors">

      </div>



      <div class="row-fluid">
        <div class="span12">

          <div class="row-fluid padbottomless">
            <h4>1. Amount</h4>
          </div>
          <div class="row-fluid">
            <div class="span12">

                <div class="radio-inline donation-v2-amounts padbottommore">
  <span>

      <input id="donation_amount_25"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="25"
              />
    <label for="donation_amount_25" class="radio">$25</label>
  </span>

  <span>

      <input id="donation_amount_50"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="50"
              />
    <label for="donation_amount_50" class="radio">$50</label>
  </span>

  <span>

      <input id="donation_amount_100"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="100"
              />
    <label for="donation_amount_100" class="radio">$100</label>
  </span>

  <span>

      <input id="donation_amount_250"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="250"
              />
    <label for="donation_amount_250" class="radio">$250</label>
  </span>

  <span>

      <input id="donation_amount_1000"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="1000"
              />
    <label for="donation_amount_1000" class="radio">$1,000</label>
  </span>

  <span>

      <input id="donation_amount_2500"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="2500"
              />
    <label for="donation_amount_2500" class="radio">$2,500</label>
  </span>

<span>
  <input id="donation_amount_other"
         type="radio"
         name="donation[amount_option]"
         class="donation_amount_option"
          />
  <label for="donation_amount_other" class="radio">Other</label>
</span>
</div>



              <div class="row-fluid donation-v2-options ">

                <div class="span6">
                  <div class="donation-other-input-container">
  <div class="currency-symbol">$</div>
  <input id="donation_amount_other_input"
         type="text"
         name="donation[amount]"
         class="text nb_donation_v2_amount nb_donation_v2_amount_$"
         />
</div>

                </div>


                <div class="span6">
                  <div class="donation-v2-occurence-radio">
                    <span>
  <input id="donation_donation_occurrence_one_time"
         type="radio"
         name="donation[donation_occurrence]"
         class="donation_amount_option"
         value="one-time"
         checked="checked" />
  <label class="radio" for="donation_donation_occurrence_one_time">
     One-time
  </label>
</span>
<span>
  <input id="donation_donation_occurrence_monthly"
        type="radio"
        name="donation[donation_occurrence]"
        class="donation_amount_option"
        value="monthly"
         />
  <label class="radio" for="donation_donation_occurrence_monthly">
    Monthly
  </label>
</span>

                  </div>
                </div>

              </div>




When you choose an amount, please remember that all donations above $25 are eligible for a charitable donation tax credit.
<br><br>
The total value of these credits varies by province or territory, ranging from a 19% to 35% refund for the first $200 donated, and a 40% - 53% refund on any amounts donated above $200.

            </div>
          </div>

          <div class="row-fluid padbottomless">
            <h4 class="sub-header">2. Your information</h4>
          </div>
          <div class="row-fluid">
            <div class="span12"><label for="donation_email">Email</label><input class="text" id="donation_email" name="donation[email]" type="email" /></div>
          </div>
          <div class="row-fluid">
            <div class="span6"><label for="donation_first_name">First Name</label><input class="text" id="donation_first_name" name="donation[first_name]" type="text" /></div>
            <div class="span6"><label for="donation_last_name">Last Name</label><input class="text" id="donation_last_name" name="donation[last_name]" type="text" /></div>

          </div>
          <div class="row-fluid">

            <div class="span12">
              <label for="donation_billing_address_country_code">Country</label><select id="donation_billing_address_country_code" name="donation[billing_address_attributes][country_code]"><option value="AF">Afghanistan</option><option value="AL">Albania</option><option value="DZ">Algeria</option><option value="AS">American Samoa</option><option value="AD">Andorra</option><option value="AO">Angola</option><option value="AI">Anguilla</option><option value="AQ">Antarctica</option><option value="AG">Antigua and Barbuda</option><option value="AR">Argentina</option><option value="AM">Armenia</option><option value="AW">Aruba</option><option value="AU">Australia</option><option value="AT">Austria</option><option value="AZ">Azerbaijan</option><option value="BS">Bahamas</option><option value="BH">Bahrain</option><option value="BD">Bangladesh</option><option value="BB">Barbados</option><option value="BY">Belarus</option><option value="BE">Belgium</option><option value="BZ">Belize</option><option value="BJ">Benin</option><option value="BM">Bermuda</option><option value="BT">Bhutan</option><option value="BO">Bolivia</option><option value="BQ">Bonaire, Sint Eustatius and Saba</option><option value="BA">Bosnia and Herzegovina</option><option value="BW">Botswana</option><option value="BV">Bouvet Island</option><option value="BR">Brazil</option><option value="IO">British Indian Ocean Territory</option><option value="BN">Brunei Darussalam</option><option value="BG">Bulgaria</option><option value="BF">Burkina Faso</option><option value="BI">Burundi</option><option value="KH">Cambodia</option><option value="CM">Cameroon</option><option value="CA" selected="selected">Canada</option><option value="CV">Cape Verde</option><option value="KY">Cayman Islands</option><option value="CF">Central African Republic</option><option value="TD">Chad</option><option value="CL">Chile</option><option value="CN">China</option><option value="CX">Christmas Island</option><option value="CC">Cocos (Keeling) Islands</option><option value="CO">Colombia</option><option value="KM">Comoros</option><option value="CG">Congo</option><option value="CD">Congo, the Democratic Republic of the</option><option value="CK">Cook Islands</option><option value="CR">Costa Rica</option><option value="HR">Croatia</option><option value="CU">Cuba</option><option value="CW">Curaçao</option><option value="CY">Cyprus</option><option value="CZ">Czech Republic</option><option value="CI">Côte d'Ivoire</option><option value="DK">Denmark</option><option value="DJ">Djibouti</option><option value="DM">Dominica</option><option value="DO">Dominican Republic</option><option value="EC">Ecuador</option><option value="EG">Egypt</option><option value="SV">El Salvador</option><option value="GQ">Equatorial Guinea</option><option value="ER">Eritrea</option><option value="EE">Estonia</option><option value="ET">Ethiopia</option><option value="FK">Falkland Islands (Malvinas)</option><option value="FO">Faroe Islands</option><option value="FJ">Fiji</option><option value="FI">Finland</option><option value="FR">France</option><option value="GF">French Guiana</option><option value="PF">French Polynesia</option><option value="TF">French Southern Territories</option><option value="GA">Gabon</option><option value="GM">Gambia</option><option value="GE">Georgia</option><option value="DE">Germany</option><option value="GH">Ghana</option><option value="GI">Gibraltar</option><option value="GR">Greece</option><option value="GL">Greenland</option><option value="GD">Grenada</option><option value="GP">Guadeloupe</option><option value="GU">Guam</option><option value="GT">Guatemala</option><option value="GG">Guernsey</option><option value="GN">Guinea</option><option value="GW">Guinea-Bissau</option><option value="GY">Guyana</option><option value="HT">Haiti</option><option value="HM">Heard Island and McDonald Islands</option><option value="VA">Holy See (Vatican City State)</option><option value="HN">Honduras</option><option value="HK">Hong Kong</option><option value="HU">Hungary</option><option value="IS">Iceland</option><option value="IN">India</option><option value="ID">Indonesia</option><option value="IR">Iran, Islamic Republic of</option><option value="IQ">Iraq</option><option value="IE">Ireland</option><option value="IM">Isle of Man</option><option value="IL">Israel</option><option value="IT">Italy</option><option value="JM">Jamaica</option><option value="JP">Japan</option><option value="JE">Jersey</option><option value="JO">Jordan</option><option value="KZ">Kazakhstan</option><option value="KE">Kenya</option><option value="KI">Kiribati</option><option value="KW">Kuwait</option><option value="KG">Kyrgyzstan</option><option value="LA">Lao People's Democratic Republic</option><option value="LV">Latvia</option><option value="LB">Lebanon</option><option value="LS">Lesotho</option><option value="LR">Liberia</option><option value="LY">Libya</option><option value="LI">Liechtenstein</option><option value="LT">Lithuania</option><option value="LU">Luxembourg</option><option value="MO">Macao</option><option value="MG">Madagascar</option><option value="MW">Malawi</option><option value="MY">Malaysia</option><option value="MV">Maldives</option><option value="ML">Mali</option><option value="MT">Malta</option><option value="MH">Marshall Islands</option><option value="MQ">Martinique</option><option value="MR">Mauritania</option><option value="MU">Mauritius</option><option value="YT">Mayotte</option><option value="MX">Mexico</option><option value="FM">Micronesia, Federated States of</option><option value="MD">Moldova, Republic of</option><option value="MC">Monaco</option><option value="MN">Mongolia</option><option value="ME">Montenegro</option><option value="MS">Montserrat</option><option value="MA">Morocco</option><option value="MZ">Mozambique</option><option value="MM">Myanmar</option><option value="NA">Namibia</option><option value="NR">Nauru</option><option value="NP">Nepal</option><option value="NL">Netherlands</option><option value="NC">New Caledonia</option><option value="NZ">New Zealand</option><option value="NI">Nicaragua</option><option value="NE">Niger</option><option value="NG">Nigeria</option><option value="NU">Niue</option><option value="NF">Norfolk Island</option><option value="KP">North Korea</option><option value="MK">North Macedonia, Republic of</option><option value="MP">Northern Mariana Islands</option><option value="NO">Norway</option><option value="OM">Oman</option><option value="PK">Pakistan</option><option value="PW">Palau</option><option value="PS">Palestine, State of</option><option value="PA">Panama</option><option value="PG">Papua New Guinea</option><option value="PY">Paraguay</option><option value="PE">Peru</option><option value="PH">Philippines</option><option value="PN">Pitcairn</option><option value="PL">Poland</option><option value="PT">Portugal</option><option value="PR">Puerto Rico</option><option value="QA">Qatar</option><option value="RO">Romania</option><option value="RU">Russian Federation</option><option value="RW">Rwanda</option><option value="RE">Réunion</option><option value="BL">Saint Barthélemy</option><option value="SH">Saint Helena, Ascension and Tristan da Cunha</option><option value="KN">Saint Kitts and Nevis</option><option value="LC">Saint Lucia</option><option value="MF">Saint Martin (French part)</option><option value="PM">Saint Pierre and Miquelon</option><option value="VC">Saint Vincent and the Grenadines</option><option value="WS">Samoa</option><option value="SM">San Marino</option><option value="ST">Sao Tome and Principe</option><option value="SA">Saudi Arabia</option><option value="SN">Senegal</option><option value="RS">Serbia</option><option value="SC">Seychelles</option><option value="SL">Sierra Leone</option><option value="SG">Singapore</option><option value="SX">Sint Maarten (Dutch part)</option><option value="SK">Slovakia</option><option value="SI">Slovenia</option><option value="SB">Solomon Islands</option><option value="SO">Somalia</option><option value="ZA">South Africa</option><option value="GS">South Georgia and the South Sandwich Islands</option><option value="KR">South Korea</option><option value="SS">South Sudan</option><option value="ES">Spain</option><option value="LK">Sri Lanka</option><option value="SD">Sudan</option><option value="SR">Suriname</option><option value="SJ">Svalbard and Jan Mayen</option><option value="SZ">Swaziland</option><option value="SE">Sweden</option><option value="CH">Switzerland</option><option value="SY">Syrian Arab Republic</option><option value="TW">Taiwan</option><option value="TJ">Tajikistan</option><option value="TZ">Tanzania, United Republic of</option><option value="TH">Thailand</option><option value="TL">Timor-Leste</option><option value="TG">Togo</option><option value="TK">Tokelau</option><option value="TO">Tonga</option><option value="TT">Trinidad and Tobago</option><option value="TN">Tunisia</option><option value="TR">Turkey</option><option value="TM">Turkmenistan</option><option value="TC">Turks and Caicos Islands</option><option value="TV">Tuvalu</option><option value="UG">Uganda</option><option value="UA">Ukraine</option><option value="AE">United Arab Emirates</option><option value="GB">United Kingdom</option><option value="US">United States</option><option value="UM">United States Minor Outlying Islands</option><option value="UY">Uruguay</option><option value="UZ">Uzbekistan</option><option value="VU">Vanuatu</option><option value="VE">Venezuela, Bolivarian Republic of</option><option value="VN">Viet Nam</option><option value="VG">Virgin Islands, British</option><option value="VI">Virgin Islands, U.S.</option><option value="WF">Wallis and Futuna</option><option value="EH">Western Sahara</option><option value="YE">Yemen</option><option value="ZM">Zambia</option><option value="ZW">Zimbabwe</option><option value="AX">Åland Islands</option></select>
            </div>

          </div>
          <div class="row-fluid">

            <div class="span12"><label for="donation_billing_address_address1">Address</label>
              <input class="text" id="donation_billing_address_address1" name="donation[billing_address_attributes][address1]" type="text" />
              <input class="text" id="donation_billing_address_address2" name="donation[billing_address_attributes][address2]" type="text" />
              <input class="text not-us-or-canada hide" id="donation_billing_address_address3" name="donation[billing_address_attributes][address3]" type="text" />
            </div>

          </div>
          <div class="row-fluid">
            <div class="span4">
              <label for="donation_billing_address_city">City/Town</label><input class="text" id="donation_billing_address_city" name="donation[billing_address_attributes][city]" type="text" />
            </div>
            <div class="span4 us-or-canada us-only hide">
              <label for="donation_state">Province</label><select id="donation_billing_address_state" name="donation[billing_address_attributes][state]"><option value="" selected="selected"></option><option value="AL">Alabama</option><option value="AK">Alaska</option><option value="AS">American Samoa</option><option value="AZ">Arizona</option><option value="AR">Arkansas</option><option value="AA">Armed Forces Americas</option><option value="AE">Armed Forces Europe</option><option value="AP">Armed Forces Pacific</option><option value="CA">California</option><option value="CO">Colorado</option><option value="CT">Connecticut</option><option value="DE">Delaware</option><option value="DC">District of Columbia</option><option value="FM">Federated States of Micronesia</option><option value="FL">Florida</option><option value="GA">Georgia</option><option value="GU">Guam</option><option value="HI">Hawaii</option><option value="ID">Idaho</option><option value="IL">Illinois</option><option value="IN">Indiana</option><option value="IA">Iowa</option><option value="KS">Kansas</option><option value="KY">Kentucky</option><option value="LA">Louisiana</option><option value="ME">Maine</option><option value="MH">Marshall Islands</option><option value="MD">Maryland</option><option value="MA">Massachusetts</option><option value="MI">Michigan</option><option value="MN">Minnesota</option><option value="MS">Mississippi</option><option value="MO">Missouri</option><option value="MT">Montana</option><option value="NE">Nebraska</option><option value="NV">Nevada</option><option value="NH">New Hampshire</option><option value="NJ">New Jersey</option><option value="NM">New Mexico</option><option value="NY">New York</option><option value="NC">North Carolina</option><option value="ND">North Dakota</option><option value="OH">Ohio</option><option value="OK">Oklahoma</option><option value="OR">Oregon</option><option value="PW">Palau</option><option value="PA">Pennsylvania</option><option value="PR">Puerto Rico</option><option value="RI">Rhode Island</option><option value="SC">South Carolina</option><option value="SD">South Dakota</option><option value="TN">Tennessee</option><option value="TX">Texas</option><option value="UT">Utah</option><option value="VT">Vermont</option><option value="VI">Virgin Island</option><option value="VA">Virginia</option><option value="WA">Washington</option><option value="WV">West Virginia</option><option value="WI">Wisconsin</option><option value="WY">Wyoming</option></select>
            </div>
            <div class="span4 us-or-canada canada-only hide">
              <label for="donation_billing_address_state">Province</label><input class="text" id="donation_billing_address_state" name="donation[billing_address_attributes][state]" type="text" />
            </div>
            <div class="span4">
              <label for="donation_billing_address_zip">Postal Code</label><input class="text" id="donation_billing_address_zip" name="donation[billing_address_attributes][zip]" type="text" />
            </div>

          </div>
          <div class="row-fluid">

            <div class="span12">
              <label for="donation_billing_address_phone_number">Phone Number</label><input class="text" id="donation_billing_address_phone_number" name="donation[billing_address_attributes][phone_number]" type="tel" />
            </div>

          </div>

          <div class="row-fluid">

            <div class="span12">
              <label class="checkbox" for="donation_email_opt_in"><input name="donation[email_opt_in]" type="hidden" value="0" /><input checked="checked" id="donation_email_opt_in" name="donation[email_opt_in]" type="checkbox" value="1" />
Send me email updates from Springtide, and allow both Springtide and Fair Voting BC to contact me for the purposes of communication about the Charter Challenge for Fair Voting Project.
<br>
                <br>
                (If you have already opted into emails from Springtide and/or the Charter Challenge project, unchecking this box will mean you will no longer receive any updates from us).
              </label>
            </div>

          </div>








          <div class="row-fluid padbottomless">
            <h4 class="sub-header">3. Payment information</h4>
          </div>
          <div class="row-fluid">
            <div class="span12">
              <label>Credit Card</label>
              <div class="payment-input payment-toggle-view" data-payments-element-type="card" data-error-container=".payment-input + div.card-error-container"></div>
<div class="card-error-container"></div>

            </div>
          </div>
          <div class="row-fluid padtopless padbottomless padtop">
            <div clas="span12">


              <label for="donation_is_private" class="checkbox padtopmore"><input name="donation[is_private]" type="hidden" value="0" /><input class="checkbox" id="donation_is_private" name="donation[is_private]" type="checkbox" value="1" /> Don't publish my donation on the website.</label>

              <span style="font-size: xx-small;">

              Note: all donations processed on this page are secure payments processed via <a href="https://stripe.com/en-ca">Stripe</a>, which has the highest level of security certification attainable in the payments industry (PCI Level 1). For alternative ways to donate, visit <a href="https://www.charterchallenge.ca/ways-to-give"> this page</a>.
              </span>
              </div>
          </div>
          <div class="row-fluid">
            <div class="span12">

            </div>
            <div class="span12">
              <div class="submit-container">
                <div class="donation-v2-amount">

                  <span>

                    <span class="hidden">$</span><span class="nb_donation_v2_amount">Please select an amount</span>


                    <div class="nb_donation_v2_interval" data-placeholder="paid monthly"></div>

                  </span>

                </div>

                <input class="submit-button btn btn-primary btn-lg" type="submit" name="commit" value="Donate now" />

              </div>
            </div>
            <div class="form-submit"></div>
          </div>
        </div>

      </div>

    </div>
  </div>

  </form>

</div>

<span style="font-size: xx-small;">
  <p> * Total reflects processed donations only. Donations made from this page are processed immediately, however some donations (e.g. cheque, e-transfer, paypal) are processed manually and there can be a delay (of up to 1 - 2 weeks) before they will be added to this total.</p>
  <p> All donations to the Charter Challenge for Fair Voting are made to the Springtide Collective for Democracy Society -  P.O. BOX 40003| Robie Street PO | Halifax, NS Canada | B3K 0E4 | 902.989.3668
    <p>Springtide is a registered Canadian charity. Our CRA charitable registration number is 838267136RR0001. For more on charitable giving in Canada, see <a href=https://www.canada.ca/charities-giving>canada.ca/charities-giving</a>
<p> <br>Contact us: <a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="3b52555d547b484b4952555c4f525f5e15555c54">[email&#160;protected]</a>
  <p>In the event that the fundraising goal is not met, Springtide (the recipient charity of all donations) retains the right to hold the funds collected for use in a future case, or alternatively, to create educational programming and resources for teaching and learning about electoral reform in Canada, or other charitable purposes as Springtide sees fit.
</span>

<script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script src="/assets/liquid/theme_donation_v2.js"></script>

  </div>

</div>
<!-- /onecolumn-container -->
<!-- /_columns_1.html -->



          </div>
          <!-- .main -->
        </div>
        <!-- .main-container -->

        <footer>
          <div class="width-container clearfix">

            <div id="fb-root"></div>

<script type="text/javascript">
  window.fbAsyncInit = function() {
    FB.init({
      appId      : 126739610711965,
      channelUrl : "//cc-springtide.nationbuilder.com/channel.html",
      status     : true,
      version    : "v9.0",
      cookie     : true,
      xfbml      : true
    });
  };
  (function(d, s, id) {
    var js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return;
    js = d.createElement(s);
    js.id = id;
    js.async = true;
    js.src = "//connect.facebook.net/en_US/sdk.js";
    fjs.parentNode.insertBefore(js, fjs);
  }(document, 'script', 'facebook-jssdk'));
</script>

<script src="https://assets.nationbuilder.com/assets/liquid-afd4cb8734a76f96f5097a424ed61c3c3354d9f9472cc52b6d1513ee749d49ec.js"></script>









            <div class="row-fluid">

              <div class="span6">

              </div>



            </div>

          </div>
          <!-- .width-container -->
        </footer>

      </div>
      <!-- #body -->
    </div>
    <!-- #wrap -->
  </div>
  <!-- #pattern -->




  <!--[if lt IE 9]>
  <script type="text/javascript" src="/javascripts/jquery.backstretch.min.js"></script>
  <script type="text/javascript">
    jQuery.backstretch("https://assets.nationbuilder.com/springtide/pages/1773/header_images/original/gavel.png?1671722376", {speed: 0});
  </script>
  <![endif]-->
  <style>
    body {
      background: url('https://assets.nationbuilder.com/springtide/pages/1773/header_images/original/gavel.png?1671722376') no-repeat center center fixed;
      -webkit-background-size: cover;
      -moz-background-size: cover;
      -o-background-size: cover;
      background-size: cover;
    }
</style>


<script type="text/javascript">
  function pollFBCommentBox(widget) {
    var fbCommentBoxTimer = window.setInterval(function() {
      if($(widget).height() < 100) {
        $(widget).removeClass('comment-box-active');
        window.clearInterval(fbCommentBoxTimer);
      }
    }, 1000);
  }
  FB.Event.subscribe('edge.create', function(href, widget) {
    $(widget).addClass('comment-box-active');
    pollFBCommentBox(widget);
  });
  FB.Event.subscribe('edge.remove', function(href, widget) {
    $(widget).removeClass('comment-box-active');
  });
</script>

  <script>
    if ( window.self !== window.top ) {
      var referrer_origin = "https://www.charterchallenge.ca";
      if ( window.location.origin !== referrer_origin ) {
        var xhttp = new XMLHttpRequest();
        var params = "iframe_req_path=" + window.location.pathname + "&referrer_origin=" + referrer_origin;
        xhttp.open("GET", "iframe_requests?" + params, true);
        xhttp.send();
      }
    }
  </script>
</body>

</html>
`

const webPageWithJustStarted = `
<!DOCTYPE html>
<!--[if lt IE 7]>
<html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
<!--[if IE 7]>
<html class="no-js lt-ie9 lt-ie8"> <![endif]-->
<!--[if IE 8]>
<html class="no-js lt-ie9"> <![endif]-->
<!--[if gt IE 8]><!-->
<html class="no-js"> <!--<![endif]-->
<head>
  <meta name="google-site-verification" content="pJgtWPNLX5_ucuNccfFyJu8LSQ-aDzrnVveSiUc23x8" />
  <meta name="google-site-verification" content="CjNDrfIadB9PJb8DiRI7IDIC2Ud_s4OdC1A9eLfNXP8" />
  <meta name="google-site-verification" content="tWVq_LM8h4MFZ51wf5z6L8SzXyA115_g_ie4MFOKijg" />
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <title>Donate Charter Challenge for Fair Voting</title>
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.css" rel="stylesheet">
  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431694346012/default/theme.scss" type="text/css" />
  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431694346012/default/tablet-and-desktop.scss" type="text/css" media="screen and (min-width: 768px)" />

  <!-- because ie8 ignores media queries, we need this -->
  <!--[if IE 8]>
    <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431694346012/default/tablet-and-desktop.scss" type="text/css" />
  <![endif]-->


  <!--[if IE]>
  <link rel="stylesheet" href="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431694346012/default/ie.scss" type="text/css" />
  <![endif]-->



<script type="text/javascript">var _sf_startpt=(new Date()).getTime()</script>



<meta content="authenticity_token" name="csrf-param" />
<meta content="/6uo995ms18WQ93rOsrCj7FwSm8FcFlF/xuXnIWOoRM=" name="csrf-token" />

  <link rel="canonical" href="https://www.charterchallenge.ca/donate_appeal" />
    <meta name="Title" content="Donate ">
    <meta property="og:title" content="Donate "/>
  <meta property="og:url" content="https://www.charterchallenge.ca/donate_appeal">
    <meta property="og:type" content="article">
      <link rel="image_src" href="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891" />
      <meta property="og:image" content="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891" />
  <meta property="og:site_name" content="Charter Challenge for Fair Voting"/>

<script type="text/javascript">
  var NB = NB || {};

  NB.environment = "production";
  NB.pageId = "1778";
  NB.Liquid = NB.Liquid || {
    Theme: {
      version: 2,
      name: "Charter Challenge for Fair Voting",
      parent: "publish",
      variation: ""
    }
  };
  NB.payments = NB.payments || {elements: {}};
  NB.payments.decimal_mark = "."
  NB.payments.thousands_separator = ","
  NB.payments.currency = "cad";
</script>

<script type="text/javascript">
    //<![CDATA[
      window._auth_token_name = "authenticity_token";
      window._auth_token = "/6uo995ms18WQ93rOsrCj7FwSm8FcFlF/xuXnIWOoRM=";
    //]]>
</script>

  <link rel="stylesheet" href="https://ajax.googleapis.com/ajax/libs/jqueryui/1.10.0/themes/cupertino/jquery-ui.css" type="text/css" media="all">

      <link rel="icon" type="image/x-icon" href="https://assets.nationbuilder.com/springtide/sites/38/favicon_images/original/faviconcc.png?1496167237" />


  <script src="https://assets.nationbuilder.com/assets/liquid/main-c2d17f5c65a7fbd197b7a65357ac82be1a4ff51b2932b32fd233152158bca307.js"></script>

<script type="text/javascript">
  window.twttr = (function (d,s,id) {
    var t, js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return; js=d.createElement(s); js.id=id;
    js.src="//platform.twitter.com/widgets.js"; fjs.parentNode.insertBefore(js, fjs);
    return window.twttr || (t = { _e: [], ready: function(f){ t._e.push(f) } });
  }(document, "script", "twitter-wjs"));
</script>

<script type="text/javascript">
  NB.FBAppId = '126739610711965';
</script>

  <script type="text/javascript">
    (function (d) { var config = { kitId: 'atg5ome', scriptTimeout: 3000, async: true }, h=d.documentElement,t=setTimeout(function(){h.className=h.className.replace(/\bwf-loading\b/g,"")+" wf-inactive";},config.scriptTimeout),tk=d.createElement("script"),f=false,s=d.getElementsByTagName("script")[0],a;h.className+=" wf-loading";tk.src='https://use.typekit.net/'+config.kitId+'.js';tk.async=true;tk.onload=tk.onreadystatechange=function(){a=this.readyState;if(f||a&&a!="complete"&&a!="loaded")return;f=true;clearTimeout(t);try{Typekit.load(config)}catch(e){}};s.parentNode.insertBefore(tk,s)})(document);
  </script>

  <script type="text/javascript">
    var _gaq = _gaq || [];
    _gaq.push(['_setAccount', 'UA-100048182-1']);
     _gaq.push(['_setDomainName', 'none']);
    _gaq.push(['_gat._anonymizeIp']);
    _gaq.push(['_setAllowLinker', true]);
      _gaq.push(['_setCustomVar', 1, 'UGC', 'false', 3]);
      _gaq.push(['_setCustomVar', 1, 'Page type', 'Donation (v2)', 3]);
    _gaq.push(['_trackPageview']);
    _gaq.push(['_trackPageLoadTime']);

    (function() {
      var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
          ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';

      let s = document.getElementsByTagName('script')[0];
      s.parentNode.insertBefore(ga, s);
    })();
  </script>






  <div id='recaptcha-input'>
    <script src="https://www.recaptcha.net/recaptcha/api.js?render=6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T"   ></script>
        <script>
          // Define function so that we can call it again later if we need to reset it
          // This executes reCAPTCHA and then calls our callback.
          function executeRecaptchaForDonate() {
            grecaptcha.ready(function() {
              grecaptcha.execute('6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T', {action: 'donate'}).then(function(token) {
                setInputWithRecaptchaResponseTokenForDonate('g-recaptcha-response-data-donate', token)
              });
            });
          };
          // Invoke immediately
          executeRecaptchaForDonate()

          // Async variant so you can await this function from another async function (no need for
          // an explicit callback function then!)
          // Returns a Promise that resolves with the response token.
          async function executeRecaptchaForDonateAsync() {
            return new Promise((resolve, reject) => {
              grecaptcha.ready(async function() {
                resolve(await grecaptcha.execute('6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T', {action: 'donate'}))
              });
            })
          };

                    var setInputWithRecaptchaResponseTokenForDonate = function(id, token) {
            var element = document.getElementById(id);
            element.value = token;
          }

        </script>
<input type="hidden" name="g-recaptcha-response-data[donate]" id="g-recaptcha-response-data-donate" data-sitekey="6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T" class="g-recaptcha g-recaptcha-response "/>

    <script>
      const clientKey = '6LekWdoZAAAAAA61Owqhdq5e8IIQEfJbEOs8IT8T'
      grecaptcha.ready( function() {
        var $badge = $('.grecaptcha-badge');
        if ($badge.length === 1){
          if ($('#cd-nav, #control-panel-nav').length === 1) {
            $badge.css("bottom", "84px")
          }
          $badge.css("z-index", "9994")
          $badge.find('iframe').css("margin", "0");
        }

        setInterval(
          function() {
            grecaptcha.execute(clientKey, { action: 'donate' }).then(
              function (token) {
                if ($("#g-recaptcha-response-data-donate").length) {
                  $("#g-recaptcha-response-data-donate").val(token);
                }
              }
            )
          }, 90000)
      });
    </script>
</div>


  <script src="https://js.stripe.com/v3/"></script>
  <script>
    NB.payments.publishableKey = "pk_live_1TjMHnI0fp51k3hKVhJEpm6D";
    NB.payments.epoButtonTheme = "dark";



    NB.payments.descriptor = "SPRINGTIDE"
  </script>
    <script src="https://assets.nationbuilder.com/assets/payments_styling-3cb98755e69cd4399888137a7648ee6eabdd63c4d4c547626e7174d769be6e1b.js"></script>
        <script>
          function initializeDefaultElementOptionsForCustomNationSignupPages() {}
        </script>
  <script src="https://assets.nationbuilder.com/assets/payments-e2d51187f660cba449c70f7f87f733f762042cf340a5fd57e0f8c1fbf885c033.js"></script>

<!-- The following line of CSS hides a checkbox in social share prompts for posting to Facebook;
As of Aug 1 2018, FB has deprecated the ability for apps to post to personal profile pages, so this is a temporary fix
while we re-configure the social share prompt -->
<style>label[for="face_tweet_is_facebook"]{display:none;}</style>


  <script type="text/javascript">
    NB.appConfig.userIsLoggedIn = false;
  </script>

  <link href='//fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic,700italic|Oswald:400,300,700|Lato:400,700,400italic,700italic|Bree+Serif|Josefin+Sans:700|Bitter:400,700,400italic|Open+Sans:400italic,700italic,400,700,800' rel='stylesheet' type='text/css'>
<link href="https://fonts.googleapis.com/css?family=Open+Sans:400,400i,700,700i" rel="stylesheet">
<link href="https://fonts.googleapis.com/css?family=Lato:900" rel="stylesheet">
<!-- Facebook Pixel Code -->
<script>
  !function(f,b,e,v,n,t,s)
  {if(f.fbq)return;n=f.fbq=function(){n.callMethod?
  n.callMethod.apply(n,arguments):n.queue.push(arguments)};
  if(!f._fbq)f._fbq=n;n.push=n;n.loaded=!0;n.version='2.0';
  n.queue=[];t=b.createElement(e);t.async=!0;
  t.src=v;s=b.getElementsByTagName(e)[0];
  s.parentNode.insertBefore(t,s)}(window, document,'script',
  'https://connect.facebook.net/en_US/fbevents.js');
  fbq('init', '119846658737416');
  fbq('track', 'PageView');
</script>
<noscript><img height="1" width="1" style="display:none" src="https://www.facebook.com/tr?id=119846658737416&ev=PageView&noscript=1" /></noscript>
<!-- End Facebook Pixel Code -->

<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431694346012/default/jquery.ui.effect.min.js"></script>
<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431694346012/default/jquery.ui.effect-slide.min.js"></script>
<script type="text/javascript" src="https://springtide.nationbuilder.com/themes/38/592ad4daed0e46ae54000000/0/attachments/14959792431694346012/default/staged-donations.js"></script>

</head>
<body class="aware-theme v2-theme page-type-donation-v2 js with-background">
 <!-- Google Tag Manager (noscript) -->
<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-P6TB763" height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
<!-- End Google Tag Manager (noscript) -->


  <div id="pattern" class="pattern">
    <div class="wrap" id="wrap">
      <div id="body" class="page-pages-show-donation-v2-wide">

        <div class="header-container clearfix">
          <div class="width-container clearfix">

            <div class="tablet-visible">

<!-- _nav.html -->


<!-- /_nav.html -->

            </div>


              <div class="site-logo">
                <header><a href="/"><img src="https://assets.nationbuilder.com/springtide/sites/38/meta_images/original/CC_LOGO_Long_Logo_-_COLOUR.png?1556419891"></a></header>
              </div>



          </div>
          <!-- .width-container -->
        </div>
        <!-- .header-container -->

        <div class="nav-container desktop-visible">
          <div class="width-container clearfix">

<!-- _nav.html -->


<!-- /_nav.html -->

          </div>
        </div>




        <div class="main-container" id="middle">
          <div class="main width-container clearfix">



<!-- _columns_1.html -->
<div class="onecolumn-container clearfix">

  <div class="columns-1-flash"><div id="flash_container">



</div>
</div>

  <div class="content-pages-show-donation-v2-wide">
    <div id="content">
  <form id="donate_appeal_page_new_donation_form" class="donation_form" method="POST" action="/forms/donations"><input name="authenticity_token" type="hidden" value="/6uo995ms18WQ93rOsrCj7FwSm8FcFlF/xuXnIWOoRM="/><input name="page_id" type="hidden" value="1778"/><input name="return_to" type="hidden" value="https://www.charterchallenge.ca/donate_appeal"/><div class="email_address_form" style="display:none;" aria-hidden="true"><p><label for="email_address">Optional email code</label><br/><input name="email_address" type="text" class="text" id="email_address" autocomplete="off"/></p></div>
  <div class="form-wrap">

      <div id="headline">
        <h2>Donate </h2>
      </div>


    <div id="intro" class="intro">
      <p>We need your support to fund the legal fees and supporting campaign for the Charter Challenge for Fair Voting - a court case against Canada's unfair voting system. </p>
<p>Over the past four years, 1200+ Canadians have collectively contributed over $230,000 to support the Notice of Application, the affidavit and evidence package, the cross-examinations, the reply evidence, and our the preparation for the court hearing. Now, we are starting to raise funds for the anticipated.</p>
<p>After the initial hearing, the costs of the case will depend on what the courts decide. If we win, government lawyers will almost certainly move to appeal that ruling. If we lose, our lawyers will move to appeal that ruling. Either way, we need to be prepared to fight for voters' rights at the appeals court.</p>
<p><span style="font-weight: 400;">You can make a secure online donation via credit card, below. For alternative ways to donate to the Charter Challenge, visit <a href="https://www.charterchallenge.ca/ways-to-give">this page</a>.</span></p>
    </div>






      <div class="clearfix">

        <div class="progress" style="width: 100%;">
          <div class="bar bar-success" style="width: 0.0%;">

            <div class="bar-text">JUST STARTED</div>

          </div>
        </div>


        <div class="bar-goal">GOAL: $23,000.00</div>

      </div>



    <div class="form">

      <div class="form-errors">

      </div>



      <div class="row-fluid">
        <div class="span12">

          <div class="row-fluid padbottomless">
            <h4>1. Amount</h4>
          </div>
          <div class="row-fluid">
            <div class="span12">

                <div class="radio-inline donation-v2-amounts padbottommore">
  <span>

      <input id="donation_amount_25"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="25"
              />
    <label for="donation_amount_25" class="radio">$25</label>
  </span>

  <span>

      <input id="donation_amount_50"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="50"
              />
    <label for="donation_amount_50" class="radio">$50</label>
  </span>

  <span>

      <input id="donation_amount_100"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="100"
              />
    <label for="donation_amount_100" class="radio">$100</label>
  </span>

  <span>

      <input id="donation_amount_250"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="250"
              />
    <label for="donation_amount_250" class="radio">$250</label>
  </span>

  <span>

      <input id="donation_amount_1000"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="1000"
              />
    <label for="donation_amount_1000" class="radio">$1,000</label>
  </span>

  <span>

      <input id="donation_amount_2500"
             type="radio"
             name="donation[amount_option]"
             class="donation_amount_option"
             value="2500"
              />
    <label for="donation_amount_2500" class="radio">$2,500</label>
  </span>

<span>
  <input id="donation_amount_other"
         type="radio"
         name="donation[amount_option]"
         class="donation_amount_option"
          />
  <label for="donation_amount_other" class="radio">Other</label>
</span>
</div>



              <div class="row-fluid donation-v2-options ">

                <div class="span6">
                  <div class="donation-other-input-container">
  <div class="currency-symbol">$</div>
  <input id="donation_amount_other_input"
         type="text"
         name="donation[amount]"
         class="text nb_donation_v2_amount nb_donation_v2_amount_$"
         />
</div>

                </div>


                <div class="span6">
                  <div class="donation-v2-occurence-radio">
                    <span>
  <input id="donation_donation_occurrence_one_time"
         type="radio"
         name="donation[donation_occurrence]"
         class="donation_amount_option"
         value="one-time"
         checked="checked" />
  <label class="radio" for="donation_donation_occurrence_one_time">
     One-time
  </label>
</span>
<span>
  <input id="donation_donation_occurrence_monthly"
        type="radio"
        name="donation[donation_occurrence]"
        class="donation_amount_option"
        value="monthly"
         />
  <label class="radio" for="donation_donation_occurrence_monthly">
    Monthly
  </label>
</span>

                  </div>
                </div>

              </div>




When you choose an amount, please remember that all donations above $25 are eligible for a charitable donation tax credit.
<br><br>
The total value of these credits varies by province or territory, ranging from a 19% to 35% refund for the first $200 donated, and a 40% - 53% refund on any amounts donated above $200.

            </div>
          </div>

          <div class="row-fluid padbottomless">
            <h4 class="sub-header">2. Your information</h4>
          </div>
          <div class="row-fluid">
            <div class="span12"><label for="donation_email">Email</label><input class="text" id="donation_email" name="donation[email]" type="email" /></div>
          </div>
          <div class="row-fluid">
            <div class="span6"><label for="donation_first_name">First Name</label><input class="text" id="donation_first_name" name="donation[first_name]" type="text" /></div>
            <div class="span6"><label for="donation_last_name">Last Name</label><input class="text" id="donation_last_name" name="donation[last_name]" type="text" /></div>

          </div>
          <div class="row-fluid">

            <div class="span12">
              <label for="donation_billing_address_country_code">Country</label><select id="donation_billing_address_country_code" name="donation[billing_address_attributes][country_code]"><option value="AF">Afghanistan</option><option value="AL">Albania</option><option value="DZ">Algeria</option><option value="AS">American Samoa</option><option value="AD">Andorra</option><option value="AO">Angola</option><option value="AI">Anguilla</option><option value="AQ">Antarctica</option><option value="AG">Antigua and Barbuda</option><option value="AR">Argentina</option><option value="AM">Armenia</option><option value="AW">Aruba</option><option value="AU">Australia</option><option value="AT">Austria</option><option value="AZ">Azerbaijan</option><option value="BS">Bahamas</option><option value="BH">Bahrain</option><option value="BD">Bangladesh</option><option value="BB">Barbados</option><option value="BY">Belarus</option><option value="BE">Belgium</option><option value="BZ">Belize</option><option value="BJ">Benin</option><option value="BM">Bermuda</option><option value="BT">Bhutan</option><option value="BO">Bolivia</option><option value="BQ">Bonaire, Sint Eustatius and Saba</option><option value="BA">Bosnia and Herzegovina</option><option value="BW">Botswana</option><option value="BV">Bouvet Island</option><option value="BR">Brazil</option><option value="IO">British Indian Ocean Territory</option><option value="BN">Brunei Darussalam</option><option value="BG">Bulgaria</option><option value="BF">Burkina Faso</option><option value="BI">Burundi</option><option value="KH">Cambodia</option><option value="CM">Cameroon</option><option value="CA" selected="selected">Canada</option><option value="CV">Cape Verde</option><option value="KY">Cayman Islands</option><option value="CF">Central African Republic</option><option value="TD">Chad</option><option value="CL">Chile</option><option value="CN">China</option><option value="CX">Christmas Island</option><option value="CC">Cocos (Keeling) Islands</option><option value="CO">Colombia</option><option value="KM">Comoros</option><option value="CG">Congo</option><option value="CD">Congo, the Democratic Republic of the</option><option value="CK">Cook Islands</option><option value="CR">Costa Rica</option><option value="HR">Croatia</option><option value="CU">Cuba</option><option value="CW">Curaçao</option><option value="CY">Cyprus</option><option value="CZ">Czech Republic</option><option value="CI">Côte d'Ivoire</option><option value="DK">Denmark</option><option value="DJ">Djibouti</option><option value="DM">Dominica</option><option value="DO">Dominican Republic</option><option value="EC">Ecuador</option><option value="EG">Egypt</option><option value="SV">El Salvador</option><option value="GQ">Equatorial Guinea</option><option value="ER">Eritrea</option><option value="EE">Estonia</option><option value="ET">Ethiopia</option><option value="FK">Falkland Islands (Malvinas)</option><option value="FO">Faroe Islands</option><option value="FJ">Fiji</option><option value="FI">Finland</option><option value="FR">France</option><option value="GF">French Guiana</option><option value="PF">French Polynesia</option><option value="TF">French Southern Territories</option><option value="GA">Gabon</option><option value="GM">Gambia</option><option value="GE">Georgia</option><option value="DE">Germany</option><option value="GH">Ghana</option><option value="GI">Gibraltar</option><option value="GR">Greece</option><option value="GL">Greenland</option><option value="GD">Grenada</option><option value="GP">Guadeloupe</option><option value="GU">Guam</option><option value="GT">Guatemala</option><option value="GG">Guernsey</option><option value="GN">Guinea</option><option value="GW">Guinea-Bissau</option><option value="GY">Guyana</option><option value="HT">Haiti</option><option value="HM">Heard Island and McDonald Islands</option><option value="VA">Holy See (Vatican City State)</option><option value="HN">Honduras</option><option value="HK">Hong Kong</option><option value="HU">Hungary</option><option value="IS">Iceland</option><option value="IN">India</option><option value="ID">Indonesia</option><option value="IR">Iran, Islamic Republic of</option><option value="IQ">Iraq</option><option value="IE">Ireland</option><option value="IM">Isle of Man</option><option value="IL">Israel</option><option value="IT">Italy</option><option value="JM">Jamaica</option><option value="JP">Japan</option><option value="JE">Jersey</option><option value="JO">Jordan</option><option value="KZ">Kazakhstan</option><option value="KE">Kenya</option><option value="KI">Kiribati</option><option value="KW">Kuwait</option><option value="KG">Kyrgyzstan</option><option value="LA">Lao People's Democratic Republic</option><option value="LV">Latvia</option><option value="LB">Lebanon</option><option value="LS">Lesotho</option><option value="LR">Liberia</option><option value="LY">Libya</option><option value="LI">Liechtenstein</option><option value="LT">Lithuania</option><option value="LU">Luxembourg</option><option value="MO">Macao</option><option value="MG">Madagascar</option><option value="MW">Malawi</option><option value="MY">Malaysia</option><option value="MV">Maldives</option><option value="ML">Mali</option><option value="MT">Malta</option><option value="MH">Marshall Islands</option><option value="MQ">Martinique</option><option value="MR">Mauritania</option><option value="MU">Mauritius</option><option value="YT">Mayotte</option><option value="MX">Mexico</option><option value="FM">Micronesia, Federated States of</option><option value="MD">Moldova, Republic of</option><option value="MC">Monaco</option><option value="MN">Mongolia</option><option value="ME">Montenegro</option><option value="MS">Montserrat</option><option value="MA">Morocco</option><option value="MZ">Mozambique</option><option value="MM">Myanmar</option><option value="NA">Namibia</option><option value="NR">Nauru</option><option value="NP">Nepal</option><option value="NL">Netherlands</option><option value="NC">New Caledonia</option><option value="NZ">New Zealand</option><option value="NI">Nicaragua</option><option value="NE">Niger</option><option value="NG">Nigeria</option><option value="NU">Niue</option><option value="NF">Norfolk Island</option><option value="KP">North Korea</option><option value="MK">North Macedonia, Republic of</option><option value="MP">Northern Mariana Islands</option><option value="NO">Norway</option><option value="OM">Oman</option><option value="PK">Pakistan</option><option value="PW">Palau</option><option value="PS">Palestine, State of</option><option value="PA">Panama</option><option value="PG">Papua New Guinea</option><option value="PY">Paraguay</option><option value="PE">Peru</option><option value="PH">Philippines</option><option value="PN">Pitcairn</option><option value="PL">Poland</option><option value="PT">Portugal</option><option value="PR">Puerto Rico</option><option value="QA">Qatar</option><option value="RO">Romania</option><option value="RU">Russian Federation</option><option value="RW">Rwanda</option><option value="RE">Réunion</option><option value="BL">Saint Barthélemy</option><option value="SH">Saint Helena, Ascension and Tristan da Cunha</option><option value="KN">Saint Kitts and Nevis</option><option value="LC">Saint Lucia</option><option value="MF">Saint Martin (French part)</option><option value="PM">Saint Pierre and Miquelon</option><option value="VC">Saint Vincent and the Grenadines</option><option value="WS">Samoa</option><option value="SM">San Marino</option><option value="ST">Sao Tome and Principe</option><option value="SA">Saudi Arabia</option><option value="SN">Senegal</option><option value="RS">Serbia</option><option value="SC">Seychelles</option><option value="SL">Sierra Leone</option><option value="SG">Singapore</option><option value="SX">Sint Maarten (Dutch part)</option><option value="SK">Slovakia</option><option value="SI">Slovenia</option><option value="SB">Solomon Islands</option><option value="SO">Somalia</option><option value="ZA">South Africa</option><option value="GS">South Georgia and the South Sandwich Islands</option><option value="KR">South Korea</option><option value="SS">South Sudan</option><option value="ES">Spain</option><option value="LK">Sri Lanka</option><option value="SD">Sudan</option><option value="SR">Suriname</option><option value="SJ">Svalbard and Jan Mayen</option><option value="SZ">Swaziland</option><option value="SE">Sweden</option><option value="CH">Switzerland</option><option value="SY">Syrian Arab Republic</option><option value="TW">Taiwan</option><option value="TJ">Tajikistan</option><option value="TZ">Tanzania, United Republic of</option><option value="TH">Thailand</option><option value="TL">Timor-Leste</option><option value="TG">Togo</option><option value="TK">Tokelau</option><option value="TO">Tonga</option><option value="TT">Trinidad and Tobago</option><option value="TN">Tunisia</option><option value="TR">Turkey</option><option value="TM">Turkmenistan</option><option value="TC">Turks and Caicos Islands</option><option value="TV">Tuvalu</option><option value="UG">Uganda</option><option value="UA">Ukraine</option><option value="AE">United Arab Emirates</option><option value="GB">United Kingdom</option><option value="US">United States</option><option value="UM">United States Minor Outlying Islands</option><option value="UY">Uruguay</option><option value="UZ">Uzbekistan</option><option value="VU">Vanuatu</option><option value="VE">Venezuela, Bolivarian Republic of</option><option value="VN">Viet Nam</option><option value="VG">Virgin Islands, British</option><option value="VI">Virgin Islands, U.S.</option><option value="WF">Wallis and Futuna</option><option value="EH">Western Sahara</option><option value="YE">Yemen</option><option value="ZM">Zambia</option><option value="ZW">Zimbabwe</option><option value="AX">Åland Islands</option></select>
            </div>

          </div>
          <div class="row-fluid">

            <div class="span12"><label for="donation_billing_address_address1">Address</label>
              <input class="text" id="donation_billing_address_address1" name="donation[billing_address_attributes][address1]" type="text" />
              <input class="text" id="donation_billing_address_address2" name="donation[billing_address_attributes][address2]" type="text" />
              <input class="text not-us-or-canada hide" id="donation_billing_address_address3" name="donation[billing_address_attributes][address3]" type="text" />
            </div>

          </div>
          <div class="row-fluid">
            <div class="span4">
              <label for="donation_billing_address_city">City/Town</label><input class="text" id="donation_billing_address_city" name="donation[billing_address_attributes][city]" type="text" />
            </div>
            <div class="span4 us-or-canada us-only hide">
              <label for="donation_state">Province</label><select id="donation_billing_address_state" name="donation[billing_address_attributes][state]"><option value="" selected="selected"></option><option value="AL">Alabama</option><option value="AK">Alaska</option><option value="AS">American Samoa</option><option value="AZ">Arizona</option><option value="AR">Arkansas</option><option value="AA">Armed Forces Americas</option><option value="AE">Armed Forces Europe</option><option value="AP">Armed Forces Pacific</option><option value="CA">California</option><option value="CO">Colorado</option><option value="CT">Connecticut</option><option value="DE">Delaware</option><option value="DC">District of Columbia</option><option value="FM">Federated States of Micronesia</option><option value="FL">Florida</option><option value="GA">Georgia</option><option value="GU">Guam</option><option value="HI">Hawaii</option><option value="ID">Idaho</option><option value="IL">Illinois</option><option value="IN">Indiana</option><option value="IA">Iowa</option><option value="KS">Kansas</option><option value="KY">Kentucky</option><option value="LA">Louisiana</option><option value="ME">Maine</option><option value="MH">Marshall Islands</option><option value="MD">Maryland</option><option value="MA">Massachusetts</option><option value="MI">Michigan</option><option value="MN">Minnesota</option><option value="MS">Mississippi</option><option value="MO">Missouri</option><option value="MT">Montana</option><option value="NE">Nebraska</option><option value="NV">Nevada</option><option value="NH">New Hampshire</option><option value="NJ">New Jersey</option><option value="NM">New Mexico</option><option value="NY">New York</option><option value="NC">North Carolina</option><option value="ND">North Dakota</option><option value="OH">Ohio</option><option value="OK">Oklahoma</option><option value="OR">Oregon</option><option value="PW">Palau</option><option value="PA">Pennsylvania</option><option value="PR">Puerto Rico</option><option value="RI">Rhode Island</option><option value="SC">South Carolina</option><option value="SD">South Dakota</option><option value="TN">Tennessee</option><option value="TX">Texas</option><option value="UT">Utah</option><option value="VT">Vermont</option><option value="VI">Virgin Island</option><option value="VA">Virginia</option><option value="WA">Washington</option><option value="WV">West Virginia</option><option value="WI">Wisconsin</option><option value="WY">Wyoming</option></select>
            </div>
            <div class="span4 us-or-canada canada-only hide">
              <label for="donation_billing_address_state">Province</label><input class="text" id="donation_billing_address_state" name="donation[billing_address_attributes][state]" type="text" />
            </div>
            <div class="span4">
              <label for="donation_billing_address_zip">Postal Code</label><input class="text" id="donation_billing_address_zip" name="donation[billing_address_attributes][zip]" type="text" />
            </div>

          </div>
          <div class="row-fluid">

            <div class="span12">
              <label for="donation_billing_address_phone_number">Phone Number</label><input class="text" id="donation_billing_address_phone_number" name="donation[billing_address_attributes][phone_number]" type="tel" />
            </div>

          </div>

          <div class="row-fluid">

            <div class="span12">
              <label class="checkbox" for="donation_email_opt_in"><input name="donation[email_opt_in]" type="hidden" value="0" /><input checked="checked" id="donation_email_opt_in" name="donation[email_opt_in]" type="checkbox" value="1" />
Send me email updates from Springtide, and allow both Springtide and Fair Voting BC to contact me for the purposes of communication about the Charter Challenge for Fair Voting Project.
<br>
                <br>
                (If you have already opted into emails from Springtide and/or the Charter Challenge project, unchecking this box will mean you will no longer receive any updates from us).
              </label>
            </div>

          </div>








          <div class="row-fluid padbottomless">
            <h4 class="sub-header">3. Payment information</h4>
          </div>
          <div class="row-fluid">
            <div class="span12">
              <label>Credit Card</label>
              <div class="payment-input payment-toggle-view" data-payments-element-type="card" data-error-container=".payment-input + div.card-error-container"></div>
<div class="card-error-container"></div>

            </div>
          </div>
          <div class="row-fluid padtopless padbottomless padtop">
            <div clas="span12">


              <label for="donation_is_private" class="checkbox padtopmore"><input name="donation[is_private]" type="hidden" value="0" /><input class="checkbox" id="donation_is_private" name="donation[is_private]" type="checkbox" value="1" /> Don't publish my donation on the website.</label>

              <span style="font-size: xx-small;">

              Note: all donations processed on this page are secure payments processed via <a href="https://stripe.com/en-ca">Stripe</a>, which has the highest level of security certification attainable in the payments industry (PCI Level 1). For alternative ways to donate, visit <a href="https://www.charterchallenge.ca/ways-to-give"> this page</a>.
              </span>
              </div>
          </div>
          <div class="row-fluid">
            <div class="span12">

            </div>
            <div class="span12">
              <div class="submit-container">
                <div class="donation-v2-amount">

                  <span>

                    <span class="hidden">$</span><span class="nb_donation_v2_amount">Please select an amount</span>


                    <div class="nb_donation_v2_interval" data-placeholder="paid monthly"></div>

                  </span>

                </div>

                <input class="submit-button btn btn-primary btn-lg" type="submit" name="commit" value="Donate now" />

              </div>
            </div>
            <div class="form-submit"></div>
          </div>
        </div>

      </div>

    </div>
  </div>

  </form>

</div>

<span style="font-size: xx-small;">
  <p> * Total reflects processed donations only. Donations made from this page are processed immediately, however some donations (e.g. cheque, e-transfer, paypal) are processed manually and there can be a delay (of up to 1 - 2 weeks) before they will be added to this total.</p>
  <p> All donations to the Charter Challenge for Fair Voting are made to the Springtide Collective for Democracy Society -  P.O. BOX 40003| Robie Street PO | Halifax, NS Canada | B3K 0E4 | 902.989.3668
    <p>Springtide is a registered Canadian charity. Our CRA charitable registration number is 838267136RR0001. For more on charitable giving in Canada, see <a href=https://www.canada.ca/charities-giving>canada.ca/charities-giving</a>
<p> <br>Contact us: <a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="630a0d050c231013110a0d04170a07064d0d040c">[email&#160;protected]</a>
  <p>In the event that the fundraising goal is not met, Springtide (the recipient charity of all donations) retains the right to hold the funds collected for use in a future case, or alternatively, to create educational programming and resources for teaching and learning about electoral reform in Canada, or other charitable purposes as Springtide sees fit.
</span>

<script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script src="/assets/liquid/theme_donation_v2.js"></script>

  </div>

</div>
<!-- /onecolumn-container -->
<!-- /_columns_1.html -->



          </div>
          <!-- .main -->
        </div>
        <!-- .main-container -->

        <footer>
          <div class="width-container clearfix">

            <div id="fb-root"></div>

<script type="text/javascript">
  window.fbAsyncInit = function() {
    FB.init({
      appId      : 126739610711965,
      channelUrl : "//cc-springtide.nationbuilder.com/channel.html",
      status     : true,
      version    : "v17.0",
      cookie     : true,
      xfbml      : true
    });
  };
  (function(d, s, id) {
    var js, fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return;
    js = d.createElement(s);
    js.id = id;
    js.async = true;
    js.src = "//connect.facebook.net/en_US/sdk.js";
    fjs.parentNode.insertBefore(js, fjs);
  }(document, 'script', 'facebook-jssdk'));
</script>

<script src="https://assets.nationbuilder.com/assets/liquid-6fda76e47cd1a46bec92e2adac0a0453c78638197e234d7667c2ff4366c5a44a.js"></script>









            <div class="row-fluid">

              <div class="span6">

              </div>



            </div>

          </div>
          <!-- .width-container -->
        </footer>

      </div>
      <!-- #body -->
    </div>
    <!-- #wrap -->
  </div>
  <!-- #pattern -->




  <!--[if lt IE 9]>
  <script type="text/javascript" src="/javascripts/jquery.backstretch.min.js"></script>
  <script type="text/javascript">
    jQuery.backstretch("https://assets.nationbuilder.com/springtide/pages/1778/header_images/original/gavel.png?1694344972", {speed: 0});
  </script>
  <![endif]-->
  <style>
    body {
      background: url('https://assets.nationbuilder.com/springtide/pages/1778/header_images/original/gavel.png?1694344972') no-repeat center center fixed;
      -webkit-background-size: cover;
      -moz-background-size: cover;
      -o-background-size: cover;
      background-size: cover;
    }
</style>


<script type="text/javascript">
  function pollFBCommentBox(widget) {
    var fbCommentBoxTimer = window.setInterval(function() {
      if($(widget).height() < 100) {
        $(widget).removeClass('comment-box-active');
        window.clearInterval(fbCommentBoxTimer);
      }
    }, 1000);
  }
  FB.Event.subscribe('edge.create', function(href, widget) {
    $(widget).addClass('comment-box-active');
    pollFBCommentBox(widget);
  });
  FB.Event.subscribe('edge.remove', function(href, widget) {
    $(widget).removeClass('comment-box-active');
  });
</script>

  <script>
    if ( window.self !== window.top ) {
      var referrer_origin = "https://www.charterchallenge.ca";
      if ( window.location.origin !== referrer_origin ) {
        var xhttp = new XMLHttpRequest();
        var params = "iframe_req_path=" + window.location.pathname + "&referrer_origin=" + referrer_origin;
        xhttp.open("GET", "iframe_requests?" + params, true);
        xhttp.send();
      }
    }
  </script>
<script defer src="https://static.cloudflareinsights.com/beacon.min.js/v8b253dfea2ab4077af8c6f58422dfbfd1689876627854" integrity="sha512-bjgnUKX4azu3dLTVtie9u6TKqgx29RBwfj3QXYt5EKfWM/9hPSAI/4qcV5NACjwAo8UtTeWefx6Zq5PHcMm7Tg==" data-cf-beacon='{"rayId":"80554ebe095836d3","token":"ff633690ee944cc4bb4fa2453466614e","version":"2023.8.0","si":100}' crossorigin="anonymous"></script>
</body>

</html>
`
