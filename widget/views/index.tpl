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

    <link rel="stylesheet" href="{{.CacheKey}}/static/css/style.css">

  </head>

  <body>
    <header>

      <h1 class="logo">Beego</h1>
      <div class="description">
        Beego is a simple & powerful Go web framework.
      </div>
      {{ componentHtml .data.Quote }}
    </header>

    <lumavate-camera></lumavate-camera>

    <footer>
      {{ componentHtml .data.NavBar }}
    </footer>
    
    <div class="backdrop"></div>

    <script type="text/javascript" src="{{.CacheKey}}/lc/lumavate-components.js"></script>
    <script type="text/javascript">
      window.cacheKey = {{.CacheKey}};

    </script>
  </body>
</html>
