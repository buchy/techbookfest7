<!DOCTYPE html>
<html lang="ja">
<head>
    <title>見せられないよ！</title>
    <meta charset="UTF-8">
</head>
<body>
<h1>見せられないよ！</h1>
<p>選択された画像から顔領域を検出して塗りつぶして表示するサービスです。</p>
<div>
    <!-- (1) -->
    {{ if .Image }}
        <img src="data:image/jpg;base64,{{ .Image }}" width="500"/>
    {{ end }}
</div>
<!-- (2) -->
<form action="/analyze" enctype="multipart/form-data" method="post">
    <input type="file" name="image" accept="image/*" required>
    <input type="submit" value="解析">
</form>
</body>
</html>