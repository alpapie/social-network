<script>

</script>

<head>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<meta http-equiv="X-UA-Compatible" content="ie=edge" />
	<title>Social network</title>
    <!-- Favicon icon -->
    <link rel="icon" type="image/png" sizes="16x16" href="images/favicon.png">
    <link rel="stylesheet" href="css/emoji.css">
    
    <link rel="stylesheet" href="css/lightbox.css">

	<link rel="stylesheet" href="css/themify-icons.css" />
	<link rel="stylesheet" href="css/feather.css" />
	<!-- Favicon icon -->

	<!-- Custom Stylesheet -->
	<link rel="stylesheet" href="css/style.css" />
	<!-- This link is for the select tag. It allow us to choose users when the post is almostprovate -->
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/habibmhamadi/multi-select-tag@2.0.1/dist/css/multi-select-tag.css">
	
</head>

<body class="color-theme-blue">
	<slot />
	<!-- <script src="js/plugin.js"></script>
		
		<script src="js/lightbox.js"></script>
		<script src="js/scripts.js"></script> -->
		<script  src="https://cdn.jsdelivr.net/gh/habibmhamadi/multi-select-tag/dist/js/multi-select-tag.js"></script>
		<script>
			new MultiSelectTag('allowed', {
    rounded: true,    // default true
    shadow: true,      // default false
    placeholder: 'Search',  // default Search...
    tagColor: {
        textColor: '#030c16',
        borderColor: '#bcd6f6',
        bgColor: '#bcd6f6',
    },
    onChange: function(values) {
        console.log(values)
    }
})
		</script>
</body>
