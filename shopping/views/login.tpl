<!DOCTYPE html>
<html class="login-alone">
    <head>
        <title>登录</title>
		<meta name="keywords" content="注册登录页面" />
		<meta name="description" content="注册登录页面" />
        <meta http-equiv="content-type" content="text/html; charset=UTF-8" />
        <link rel="shortcut icon" type="image/x-icon" href="res/homepage/favicon.ico?v=3.9" />
        <link rel="stylesheet" type="text/css" href="static/css/screen.css?v=3.9" media="screen, projection">
        <link rel="stylesheet" type="text/css" href="static/css/base.css?v=3.9">
        <link rel="stylesheet" type="text/css" href="static/css/login.css?v=3.9">
       
    </head>
    <body>
        <div class="logina-logo" style="height: 55px">
            <a href="http://www.js-css.cn">
                商城
            </a>
        </div>
        <div class="logina-main main clearfix">
            <div class="tab-con">
                <form id="form-login" method="post" action="http://www.js-css.cn">
                    <div id='login-error' class="error-tip"></div>
                    <table border="0" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <th>账户</th>
                                <td width="245">
                                    <input id="email" type="text" name="email" placeholder="电子邮箱/手机号" autocomplete="off" value=""></td>
                                <td>
                                </td>
                            </tr>
                            <tr>
                                <th>密码</th>
                                <td width="245">
                                    <input id="password" type="password" name="password" placeholder="请输入密码" autocomplete="off">
                                </td>
                                <td>
                                </td>
                            </tr>
                            <tr class="find">
                                <th></th>
                                <td>
                                    <div>
                                        <label class="checkbox" for="chk11"><input style="height: auto;" id="chk11" type="checkbox" name="remember_me" >记住我</label>
                                        <a href="#">忘记密码？</a>
                                    </div>
                                </td>
                                <td></td>
                            </tr>
                            <tr>
                                <th></th>
                                <td width="245"><input class="confirm" type="submit" value="登  录"></td>
                                <td></td>
                            </tr>
                        </tbody>
                    </table>
                    <input type="hidden" name="refer" value="site/">
                </form>
            </div>
            <div class="reg">
                <p>还没有账号？<br>赶快免费注册一个吧！</p>
                <a class="reg-btn" href="#">立即免费注册</a>
            </div>
        </div>
        <div id="footer">
            <div class="copyright">版权所有</div>
        </div>
            
    </body>
</html>
