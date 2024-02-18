<script>
   
  import { page } from '$app/stores';
</script>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{$page.status}</title>

    <link rel="stylesheet" href="/css/themify-icons.css">
    <link rel="stylesheet" href="/css/feather.css">
    <!-- Favicon icon -->
    <link rel="icon" type="image/png" sizes="16x16" href="/images/favicon.png">
    <!-- Custom Stylesheet -->
    <link rel="stylesheet" href="css/style.css">

</head>

<body class="color-theme-blue mont-font">

    <!-- <div class="preloader"></div> -->

    
    <div class="main-wrapper">


        <!-- main content -->
        <div class="main-content pt-0 bg-white ps-0 pe-0">
            <div class="container">
                <div class="row justify-content-center">
                    <div class="col-lg-6 col-md-8 text-center default-page vh-100 align-items-center d-flex">
                        <div class="card border-0 text-center d-block p-0">
                            <img src="images/bg-43.png" alt="icon" class="w200 mb-4 ms-auto me-auto pt-md-5">
                            <h1 class="fw-700 text-grey-900 display3-size display4-md-size">Oops!error {$page.status}</h1>
                            <p class="text-grey-500 font-xsss">{$page.error.message} </p>
                            <a href="/" class="p-3 w175 bg-current text-white d-inline-block text-center fw-600 font-xssss rounded-3 text-uppercase ls-3">Home Page</a>
                        </div>
                    </div>
                </div>
            </div> 
        </div>
        <!-- main content -->

        
        <div class="app-footer border-0 shadow-lg bg-primary-gradiant">
            <a href="default.html" class="nav-content-bttn nav-center"><i class="feather-home"></i></a>
            <a href="default-follower.html" class="nav-content-bttn"><i class="feather-package"></i></a>
            <a href="default-live-stream.html" class="nav-content-bttn" data-tab="chats"><i class="feather-layout"></i></a>            
            <a href="shop-2.html" class="nav-content-bttn"><i class="feather-layers"></i></a>
            <a href="default-settings.html" class="nav-content-bttn"><img src="images/female-profile.png" alt="user" class="w30 shadow-xss"></a>
        </div>

        <div class="app-header-search">
            <form class="search-form">
                <div class="form-group searchbox mb-0 border-0 p-1">
                    <input type="text" class="form-control border-0" placeholder="Search...">
                    <i class="input-icon">
                        <ion-icon name="search-outline" role="img" class="md hydrated" aria-label="search outline"></ion-icon>
                    </i>
                    <a href="/" class="ms-1 mt-1 d-inline-block close searchbox-close">
                        <i class="ti-close font-xs"></i>
                    </a>
                </div>
            </form>
        </div>

    </div>
    
    <script src="js/plugin.js"></script>
    <script src="js/countdown.js"></script> 
    <script src="js/scripts.js"></script>

</body>

</html>