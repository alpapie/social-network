<script>
import { enhance } from "$app/forms";
import Login_header from "../login-header.svelte";

export let form 
</script>
<style>
	.pbg-dark{
		color: brown  !important;; 
	}
</style>

<!-- <div class="preloader"></div> -->

<div class="main-wrap">
  <Login_header />

  <div class="row">
    <div
      class="col-xl-5 d-none d-xl-block p-0 vh-100 bg-image-cover bg-no-repeat"
      style="background-image: url(images/login-bg.jpg);"
    ></div>
    <div
      class="col-xl-7 vh-100 align-items-center d-flex bg-white rounded-3 overflow-hidden"
    >
	{#if form}
		{form.error}
	{/if}
      <div class="card shadow-none border-0 ms-auto me-auto login-card">
        <div class="card-body rounded-0 text-left">
          <h2 class="fw-700 display1-size display2-md-size mb-3">
            Login into <br />your account
          </h2>
          <form method="POST" use:enhance >
            <div class="form-group icon-input mb-3">
              <i class="font-sm ti-email text-grey-500 pe-0"></i>
              <input required
                type="text"
                class="style2-input ps-5 form-control text-grey-900 font-xsss fw-600"
                placeholder="Your Email Address" name="email"
              />
            </div>
            <div class="form-group icon-input mb-1">
              <input
                type="Password" required
                class="style2-input ps-5 form-control text-grey-900 font-xss ls-3"
                placeholder="Password" name="password"
              />
              <i class="font-sm ti-lock text-grey-500 pe-0"></i>
            </div>
            <div class="form-group mb-1">
              <button t class="form-control text-center style2-input text-white fw-600 bg-dark border-0 p-0">Login</button>
            </div>
          </form>
          <br />
          <div class="col-sm-12 p-0 text-left">
            <h6 class="text-grey-500 font-xsss fw-500 mt-0 mb-0 lh-32">
              Dont have account <a href="/register" class="fw-700 ms-1"
                >Register</a
              >
            </h6>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
