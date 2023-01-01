## 云村音乐机器人

网易云音乐机器人，Telegram 搜索 @ncloud_music_bot 

### 功能
- [x] 关注云村用户，
- [x] 查看关注用户听歌榜
- [x] 音乐搜索、下载、推送
- [ ] 关注用户歌单更新推送
- [ ] 关注人有效性验证， 取消关注

![img.png](img.png)

### 运行
```
docker pull nautilis/netease_music_bot
docker run -tid -e NetEaseAccount="your neteaset account" -e NetEasePwd="md5(your netease password)" -e TelegramToken="your telegram bot token" nautilis/netease_music_bot:latest 
```