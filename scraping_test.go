package main

import "testing"

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
