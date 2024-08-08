// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import viewCommon "github.com/datamonsterr/PTsolana/view/common"

func Index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"EnVizion Gym Front Page\"><meta name=\"keywords\" content=\"front page\"><title>Gym</title><link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css\" integrity=\"sha512-DTOQO9RWCH3ppGqcWaEA1BIZOC6xxalwEsw9c2QQeAIftl+Vegovlnee1c9QX4TctnWMn13TZye+giMm8e2LwA==\" crossorigin=\"anonymous\" referrerpolicy=\"no-referrer\"><link rel=\"stylesheet\" href=\"./styles/general.css\"><link rel=\"stylesheet\" href=\"./styles/header.css\"><link rel=\"stylesheet\" href=\"./styles/footer.css\"><link rel=\"stylesheet\" href=\"./styles/index.css\"><link rel=\"stylesheet\" href=\"./styles/user-sign.css\"><link rel=\"stylesheet\" href=\"./styles/animations.css\"><link rel=\"icon\" type=\"image/x-icon\" href=\"./logos/sparta-logo-red.png\"></head><body><!-- Header Navbar Section --><header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = viewCommon.Navbar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</header><main><!-- User Login || Sign In  --><div class=\"user-sign-flex\"><section class=\"user-sign-in-container\"><span class=\"icon-close\" id=\"closeButton\"><ion-icon name=\"close-circle\"></ion-icon></span><div class=\"form-box login\"><h2>Login</h2><form action=\"#\" method=\"GET\"><div class=\"input-box\"><span class=\"icon\"><ion-icon name=\"mail\"></ion-icon></span> <input type=\"email\" required> <label for=\"email\">Email</label></div><div class=\"input-box\"><span class=\"icon\"><ion-icon name=\"lock-closed\"></ion-icon></span> <input type=\"password\" required> <label for=\"password\">Password</label></div><div class=\"remember-forget-pass\"><label><input type=\"checkbox\">Remember Me</label> <a href=\"#\" class=\"forgot-pass\">Forgot Password?</a></div><button type=\"submit\" class=\"submit-btn\">Login</button><div class=\"login-register\"><p>Don't have an Account?<a href=\"#\" class=\"register-link\">Register</a></p></div></form></div><span class=\"icon-close\" id=\"closeButton\"><ion-icon name=\"close\"></ion-icon></span><div class=\"form-box register\"><h2>Register</h2><form action=\"#\" method=\"GET\"><div class=\"input-box\"><span class=\"icon\"><ion-icon name=\"person\"></ion-icon></span> <input type=\"text\" required> <label for=\"text\">Username</label></div><div class=\"input-box\"><span class=\"icon\"><ion-icon name=\"mail\"></ion-icon></span> <input type=\"email\" required> <label for=\"email\">Email</label></div><div class=\"input-box\"><span class=\"icon\"><ion-icon name=\"lock-closed\"></ion-icon></span> <input type=\"password\" required> <label for=\"password\">Password</label></div><div class=\"remember-forget-pass\"><label><input type=\"checkbox\">I agree to the terms & conditions</label></div><button type=\"submit\" class=\"submit-btn\">Sign In</button><div class=\"login-register\"><p>Already have an Account?<a href=\"#\" class=\"login-link\">Login</a></p></div></form></div></section></div><!-- Hero Section --><section class=\"hero\"><div class=\"hero-content\"><h1 class=\"hero-title-animate\">Clarity in <span class=\"delay-span-animate\">Vision</span>, Strength in <span class=\"delay-span-animate\">Action</span></h1><h1 class=\"hero-title-animate-2\">Envizion Your Future</h1><p class=\"hero-animate-p\">Join the Movement!</p><a href=\"./membership.html\" class=\"get-started-btn gsap-animate-btn\">Get Started</a></div></section><!-- Services Section --><section class=\"services\"><div class=\"services-container\"><div><div class=\"hiddenLeft\"><p class=\"topline\">Features</p><h1 class=\"services-heading\">What We Offer </h1></div><div class=\"services-features hiddenLeft\"><p class=\"services-feature delay hiddenLeft\"><i class=\"fa-solid fa-clock services-icon\"></i> 24/7 Access</p><p class=\"services-feature delay hiddenLeft\"><i class=\"fa-solid fa-user-group services-icon\"></i> Personal Training Sessions</p><p class=\"services-feature delay hiddenLeft\"><i class=\"fa-solid fa-hot-tub-person services-icon\"></i> Steam Room and Sauna</p><p class=\"services-feature delay hiddenLeft\"><i class=\"fa-solid fa-dumbbell services-icon\"></i> Top of the line equipment</p><p class=\"services-feature delay hiddenLeft\"><i class=\"fa-solid fa-basketball services-icon\"></i> Basketball Court</p><p class=\"services-feature delay hiddenLeft\"><i class=\"fa-solid fa-table-tennis-paddle-ball services-icon\"></i> Table-Tennis Playground</p><p class=\"services-feature delay hiddenLeft\"><i class=\"fa-solid fa-person-swimming services-icon\"></i> Indoor Pool</p></div></div><div><img src=\"./images/hero-background-2.jpg\" alt=\"\" class=\"services-img img-animate-top-right\"></div><div><img src=\"./images/hero-background-4.jpg\" alt=\"\" class=\"services-img img-animate-bottom-left\"></div><div><img src=\"./images/hero-background-3.jpg\" alt=\"\" class=\"services-img img-animate-bottom-right\"></div></div></section><!-- Personal Trainers Section --><section class=\"team\"><div class=\"team-wrapper\"><div class=\"team-text animate-team\"><p class=\"topline\">Private Coaching</p><h1>Meet our Trainers</h1><p class=\"team-desc\">All our hired trainers have 25+ years of experience combined. Each trainer is specialized in strength and mobility training </p></div><div class=\"team-card animate-team\"><p>Victoria</p><img src=\"./images/trainer-1.jpg\" alt=\"trainer-1\" class=\"team-img\" title=\"Victoria | Personal Trainer\"></div><div class=\"team-card animate-team\"><p>Michael</p><img src=\"./images/trainer-3.jpg\" alt=\"trainer-3\" class=\"team-img\" title=\"Michael | Personal Trainer\"></div><div class=\"team-card animate-team\"><p>Lucia</p><img src=\"./images/trainer-2.jpg\" alt=\"trainer-2\" class=\"team-img\" title=\"Lucia | Personal Trainer\"></div><div class=\"team-card animate-team\"><p><span class=\"color-change\">Oliver</span></p><img src=\"./images/trainer-4.jpg\" alt=\"trainer-4\" class=\"team-img\" title=\"Oliver | Personal Trainer\"></div><div class=\"team-card animate-team\"><p>Marko</p><img src=\"./images/trainer-5.jpg\" alt=\"trainer-5\" class=\"team-img\" title=\"Marko | Personal Trainer\"></div><div class=\"team-card animate-team\"><p>Anna</p><img src=\"./images/trainer-6.jpg\" alt=\"trainer-6\" class=\"team-img\" title=\"Anna | Personal Trainer\"></div></div></section><!-- Email Newsletter Section --><section class=\"email\"><div class=\"email-content\"><h1 class=\"animate-email\">Get Access to Members Only Perks</h1><p class=\"animate-email\">Sign up for our newsnetter belowe to get 50% off your first personal training session!</p><form action=\"#\" method=\"GET\" class=\"animate-email\"><div class=\"form-wrap\"><label for=\"email\"><input type=\"email\" placeholder=\"Enter your email\" id=\"email\" required></label> <button class=\"get-started-btn email-btn\" type=\"submit\" for=\"email\">Sign up</button></div></form></div></section></main><!-- Footer Section --><footer class=\"footer\"><div class=\"footer-wrapper\"><div class=\"footer-desc\"><h1>En<span>Vizion</span>Gym</h1><p><a href=\"mailto:envizion@gym.com\" class=\"footer-email\"><i class=\"fa-solid fa-envelope footer-icon\"></i>envizion@gym.com</a></p><p id=\"phone\"><i class=\"fa-solid fa-phone footer-icon\"></i><a href=\"tel:347- 885 - 2503\">347- 885 - 2503</a></p><p><i class=\"fa-solid fa-location-dot footer-icon\"></i>Brooklyn, New York</p></div><div class=\"footer-links\"><h2 class=\"footer-title\">Contact us</h2><a href=\"./contact.html\" class=\"footer-link\">Contact</a> <a href=\"./contact.html\" class=\"footer-link\">Sponsorships</a></div><div class=\"footer-links\"><h2 class=\"footer-title\">Memberships</h2><a href=\"./membership.html\" class=\"footer-link\">Pricing</a> <a href=\"./membership.html\" class=\"footer-link\">Plans</a> <a href=\"./membership.html\" class=\"footer-link\">FAQ</a></div><div class=\"footer-links socials\"><h2 class=\"footer-title\">Socials</h2><div><i class=\"fa-brands fa-instagram footer-icon\"></i> <a href=\"#\" class=\"footer-link\">Instagram</a></div><div><i class=\"fa-brands fa-facebook footer-icon\"></i> <a href=\"#\" class=\"footer-link\">Facebook</a></div><div><i class=\"fa-brands fa-tiktok footer-icon\"></i> <a href=\"#\" class=\"footer-link\">TikTok</a></div><div><i class=\"fa-brands fa-youtube footer-icon \"></i> <a href=\"#\" class=\"footer-link\">YouTube</a></div><div><i class=\"fa-brands fa-x-twitter footer-icon\"></i> <a href=\"#\" class=\"footer-link\">Twitter</a></div></div><div class=\"footer-links info\"><h2 class=\"footer-title\">Informations</h2><p><a href=\"#\" class=\"footer-link\">Condition of Sales</a></p><p><a href=\"#\" class=\"footer-link\">Payment security statement</a></p></div><div class=\"footer-links payment\"><div><a href=\"https://usa.visa.com/\" target=\"_blank\"><img src=\"./images/visa-card.png\" alt=\"visa-card\" class=\"footer-img\"></a></div><a href=\"https://www.paypal.com/\" target=\"_blank\"><img src=\"./images/paypal-seeklogo.com.svg\" alt=\"visa-card\" class=\"footer-img paypal\"></a></div><div class=\"footer-links payment\"><div><a href=\"https://www.mastercard.us/en-us.html\" target=\"_blank\"><img src=\"./images/mastercard-logo.svg\" alt=\"visa-card\" class=\"footer-img mastercard\"></a></div></div></div></footer><div class=\"copyright\"><p>Copyright &copy datamonsterr</p><a href=\"https://github.com/datamonsterr\"><i class=\"fa-brands fa-github cursor\"></i></a></div></body><script src=\"./js/mobile-menu.js\"></script><script src=\"./js/user-sign-in.js\"></script><script type=\"module\" src=\"https://unpkg.com/ionicons@7.1.0/dist/ionicons/ionicons.esm.js\"></script><script nomodule src=\"https://unpkg.com/ionicons@7.1.0/dist/ionicons/ionicons.js\"></script><script defer src=\"./js/animation.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/gsap.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/Flip.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/ScrollTrigger.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/Observer.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/ScrollToPlugin.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/Draggable.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/MotionPathPlugin.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/EaselPlugin.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/PixiPlugin.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/TextPlugin.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/EasePack.min.js\"></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.4/CustomEase.min.js\"></script></html>`")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
