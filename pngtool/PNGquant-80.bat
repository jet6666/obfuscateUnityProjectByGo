echo "pngquant compress..."
for /R %%i in (*.png) do  pngquant  128 --quality 80 -f --ext .png   "%%i"
pause