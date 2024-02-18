<script>
</script>

<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta http-equiv="X-UA-Compatible" content="ie=edge" />
  <title>Social network</title>
  <!-- Favicon icon -->
  <link rel="icon" type="image/png" sizes="16x16" href="/images/favicon.png" />
</head>

<body class="color-theme-blue">
  <slot />

  <script
    src="https://cdn.jsdelivr.net/gh/habibmhamadi/multi-select-tag/dist/js/multi-select-tag.js"
  ></script>
  <script>
    new MultiSelectTag("allowed", {
      rounded: true, // default true
      shadow: true, // default false
      placeholder: "Search", // default Search...
      tagColor: {
        textColor: "#030c16",
        borderColor: "#bcd6f6",
        bgColor: "#bcd6f6",
      },
      onChange: function (values) {
        console.log(values);
      },
    });
  </script>
</body>
