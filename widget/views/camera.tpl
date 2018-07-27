<!DOCTYPE html>

<html lang="en">
  <head>
    <title>
        {{.data.InstanceName}}
    </title>

    <meta charset="utf-8">
    <base href="{{.WidgetUrlPrefix}}">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <meta name="description" content="Basic Go Widget">
    <meta name="theme-color" content="#ffffff">
    <noscript>Javascript required for this site to work.</noscript>

    <link rel="apple-touch-icon" sizes="180x180" href="/iot/favicon-180x180.png">
    <link rel="icon" type="image/png" sizes="32x32" href="/iot/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/iot/favicon-16x16.png">
    <link rel="manifest" href="/manifest.json">
    <link rel="shortcut icon" href="/iot/favicon.ico">

    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">

    <link rel="stylesheet" href="{{.CacheKey}}/static/css/camera.css">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons"
      rel="stylesheet">

  </head>

  <body>
    <header>
        <lumavate-header font-color="white" background-color="#23516A" show-back-button=true Text="Camera"></lumavate-header>
    </header>

    <lumavate-camera></lumavate-camera>

    <footer>
    </footer>

    <script type="text/javascript" src="{{.CacheKey}}/lc/lumavate-components.js"></script>
    <script type="text/javascript">
      window.cacheKey = {{.CacheKey}};

    </script>
  </body>
</html>