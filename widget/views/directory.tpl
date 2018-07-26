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

    <link rel="stylesheet" href="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css">
    <link rel="stylesheet" href="{{.CacheKey}}/static/css/directory.css">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons"
      rel="stylesheet">
  </head>

  <body>
    <header>
        <lumavate-header font-color="white" background-color="#23516A" show-back-button=true Text="Your Gallery"></lumavate-header>
    </header>

    {{ if .visible }}
        <div class="outer">
            {{ range .images}}
                <div class="border">
                    <div class="inner">
                            <div class="allImages"><img src="{{.SecureUrl}}"></div>
                    </div>
                </div>
            {{end}}
        </div>

        {{ range .albums }}
            <div class="albumBorder">
                <div class="albumEdge">
                    <div class="header">
                        <p style="float: left; font-size: 13pt;">{{ .AlbumTitle }}</p>
                        <p style="float: right; color: #707072;">{{ .PictureCount }} Picture(s)</p>
                    </div>
                    <div style="clear: both;"></div>
                    <div class="images">
                        {{ if .AlbumImages }}
                            {{range .AlbumImages}}
                                <img style="width: 50px;" src="{{.SecureUrl}}">
                            {{end}}
                        {{ else }}
                            <div style="margin: 55px;"></div>
                        {{ end }}
                    </div>
                    <div class="addDelete">
                      <button onclick="window.location = '{{ $.WidgetInstancePrefix }}/album?album={{ .AlbumTitle }}'" class="mdc-button">
                          <i class="material-icons mdc-button__icon">photo_album</i>
                      </button>
                      <button onclick="window.location = '{{ $.WidgetInstancePrefix }}/add?album={{ .AlbumTitle }}'" class="mdc-button">
                          <i class="material-icons mdc-button__icon">add</i>
                      </button>
                      <button class="mdc-button">
                          <i class="material-icons mdc-button__icon">delete</i>
                      </button>
                    </div>
                </div>
            </div>
        {{ end }}

    {{ else }}
        <div id="none">
            <div>You have no albums yet.</div>
        </div>
    {{ end }}

    <div id="album" style="display: none">
        <div class="newAlbum">
            <form method="post">
                <input type="text" id="albumTitle" name="albumTitle" placeholder="Album Title" class="mdc-text-field__input">
                <p style="float: right; color: #707072;">0 Picture(s)</p>
                <div style="clear: both;"></div>
                <div class="buttons">
                    <button type="submit" class="mdc-button">
                        <i class="material-icons mdc-button__icon">save</i>
                    </button>
                    <button onclick="deleteAlbum()" class="mdc-button">
                        <i class="material-icons mdc-button__icon">delete</i>
                    </button>
                </div>
            </form>
        </div>
    </div>

    <div style="padding: 35px;"></div>

    <div class="camera">
        <button class="mdc-fab" onclick='window.location = "{{ .WidgetInstancePrefix }}/camera"' aria-label="Camera">
            <span class="mdc-fab__icon material-icons">photo_camera</span>
        </button>
    </div>

    <div class="add">
        <button onclick="addAlbum()" class="mdc-fab" aria-label="Add">
            <span class="mdc-fab__icon material-icons">add</span>
        </button>
    </div>

    <footer>

    </footer>

    <script type="text/javascript" src="{{.CacheKey}}/lc/lumavate-components.js"></script>
    <script
        src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
        integrity="sha256-3edrmyuQ0w65f8gfBsqowzjJe2iM6n0nKciPUp8y+7E="
        crossorigin="anonymous"></script>
    <script type="text/javascript">
      window.cacheKey = {{.CacheKey}};

      var albumHighlight = 'selectImage';
      var $albums = $('.albumEdge').on("click", function(e) {
        $(this).toggleClass(albumHighlight);
      })

      function addAlbum() {
        document.getElementById('album').style.display ="block";
        document.getElementById('none').style.display = "none";
      }

      function deleteAlbum() {
        document.getElementById('album').style.display = "none";
      }
    </script>
  </body>
</html>