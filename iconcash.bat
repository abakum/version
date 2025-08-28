taskkill /f /im explorer.exe
cd /d %userprofile%\AppData\Local
attrib -h IconCache.db
del IconCache.db
cd /d %userprofile%\AppData\Local\Microsoft\Windows\Explorer
attrib -h iconcache_*.db
del iconcache_*.db
start explorer.exe
