<script>
    import { enhance } from "$app/forms";
	// @ts-nocheck
    import Login_header from '../login-header.svelte'
</script>
<style>
    .form{
        display: flex;
    }
</style>
<div class="main-wrap">

    <Login_header/>

    <div class="row">
        <div class="col-xl-5 d-none d-xl-block p-0 vh-100 bg-image-cover bg-no-repeat" style="background-image: url(images/login-bg-2.jpg);"></div>
        <div class="col-xl-7 vh-100 align-items-center d-flex bg-white rounded-3 overflow-hidden">
            <div class="card shadow-none border-0 ms-auto me-auto login-card">
                <div class="card-body rounded-0 text-left">
                    <h2 class="fw-700 display1-size display2-md-size mb-4">Create <br>your account</h2>                        
                    <form method="POST" use:enhance enctype="multipart/form-data">
                        <div class="form">
                        <div class="left-column" style="flex: 1; padding-right: 10px;">
                        <div class="form-group icon-input mb-3">
                            <i class="font-sm ti-email text-grey-500 pe-0"></i>
                            <input type="text" class="style2-input ps-5 form-control text-grey-900 font-xsss fw-600" placeholder="Your Email Address" name="email">                        
                        </div>
                        <div class="form-group icon-input mb-3">
                            <input type="Password" class="style2-input ps-5 form-control text-grey-900 font-xss ls-3" placeholder="Password" name="password">
                            <i class="font-sm ti-lock text-grey-500 pe-0"></i>
                        </div>
                        <div class="form-group icon-input mb-3">
                            <i class="font-sm ti-user text-grey-500 pe-0"></i>
                            <input type="text" class="style2-input ps-5 form-control text-grey-900 font-xsss fw-600" placeholder="First Name" name="firstname">                        
                        </div>
                        <div class="form-group icon-input mb-3">
                            <i class="font-sm ti-user text-grey-500 pe-0"></i>
                            <input type="text" class="style2-input ps-5 form-control text-grey-900 font-xsss fw-600" placeholder="Last Name" name="lastname">                        
                        </div>
                    </div>
                        
                    <div class="right-column" style="flex: 1; padding-left: 10px;">
                        <div class="form-group icon-input mb-3">
                            <i class="font-sm ti-calendar text-grey-500 pe-0"></i>
                            <input type="date" class="style2-input ps-5 form-control text-grey-900 font-xsss fw-600" name="birthdate">
                        </div>
                        <div class="form-group icon-input mb-3">
                            <i class="font-sm ti-upload text-grey-500 pe-0"></i>
                            <input type="file" class="style2-input ps-5 form-control text-grey-900 font-xsss fw-600" name="avatar">
                        </div>
                        <div class="form-group icon-input mb-3">
                            <i class="font-sm ti-user text-grey-500 pe-0"></i>
                            <input type="text" class="style2-input ps-5 form-control text-grey-900 font-xsss fw-600" placeholder="Nickname" name="nickname">                        
                        </div>
                        <div class="form-group icon-input mb-3">
                            <i class="font-sm ti-comment text-grey-500 pe-0"></i>
                            <textarea class="style2-input ps-5 form-control text-grey-900 font-xsss fw-600" placeholder="Bio" style="line-height: 1.5;" name="bio"></textarea>
                        </div>
                    </div>
                </div>
                    <div class="form-group mb-1">
                        <button t class="form-control text-center style2-input text-white fw-600 bg-dark border-0 p-0">Register</button>
                    </div>  
                    </form>
                    
                     <br>
                    <div class="col-sm-12 p-0 text-left">
                        <h6 class="text-grey-500 font-xsss fw-500 mt-0 mb-0 lh-32">Already have account <a href="/login" class="fw-700 ms-1">Login</a></h6>
                    </div>
                     
                </div>
            </div> 
        </div>
    </div>
</div>
