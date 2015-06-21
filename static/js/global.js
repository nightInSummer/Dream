//base64解码方法
function base64_decode(e){var t="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",n,r,i,s,o,u,a,f,l=0,c=0,h="",p=[];if(!e)return e;e+="";do s=t.indexOf(e.charAt(l++)),o=t.indexOf(e.charAt(l++)),u=t.indexOf(e.charAt(l++)),a=t.indexOf(e.charAt(l++)),f=s<<18|o<<12|u<<6|a,n=f>>16&255,r=f>>8&255,i=f&255,u==64?p[c++]=String.fromCharCode(n):a==64?p[c++]=String.fromCharCode(n,r):p[c++]=String.fromCharCode(n,r,i);while(l<e.length);return h=p.join(""),h.replace(/\0+$/,"")}
//获取url参数
function GetQueryString(name) { var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)"); var r = window.location.search.substr(1).match(reg); if (r!=null) return unescape(r[2]); return null; }
//数组随机取数方法
function getArrayItems(arr,num){var temp_array=new Array();for(var index in arr){temp_array.push(arr[index])}var return_array=new Array();for(var i=0;i<num;i++){if(temp_array.length>0){var arrIndex=Math.floor(Math.random()*temp_array.length);return_array[i]=temp_array[arrIndex];temp_array.splice(arrIndex,1)}else{break}}return return_array};
//刷新导航条
function refreshNav(){
	if (cookieExist("WOKUID")) {
		$(".m-nav #j-login").hide();
		$(".m-nav #j-register").hide();
		$(".m-nav .j-logined #j-name").text($.cookie("WOKUNAME"));
		$(".m-nav .j-logined").show();
		var uploadSize = $.cookie("WOKUUPSIZE");
		var allUploadSize = $.cookie("WOKUALLUPSIZE");
		$(".m-nav #j-upload").text((uploadSize / 1024 / 1024).toFixed(1) + "M");
		$(".m-nav #j-upload-progress").css("width", (uploadSize / 1024 / 1024) / allUploadSize * 100 + "%");
		var free = $.cookie("WOKUFREE");
		$(".m-nav #j-money").text(free + "元");
		$(".m-nav #j-money-progress").css("width", (free / 2000) + "%");
		var read = parseInt($.cookie("WOKUREAD"));
		if (read > 0){
			if (read <= 9){
				$(".m-nav .info-number").text(read);
			}else{
				$(".m-nav .info-number").text("9+");
			}
			$(".m-nav .info-number").show();
		}
	} else {
		$(".m-nav #j-login").show();
		$(".m-nav #j-register").show();
		$(".m-nav .j-logined").hide();
	}
}
//弹出提示框
var alertInfoTimeOut;
function showAlert(content, success) {
	$(".m-layer #j-layer-content").text(content);
	var time = 500;
	if (success) {
		$(".m-layer .lymask").hide();
		$(".m-layer .lywrap").css("border-color","#5CB85C");
	} else {
		time = 1000;
		$(".m-layer .lymask").show();
		$(".m-layer .lywrap").css("border-color","#C9302C");
	}
	$(".m-layer").fadeIn(500);
	clearTimeout(alertInfoTimeOut);
	alertInfoTimeOut = setTimeout(function() {
		$(".m-layer").fadeOut(500)
	},
	time)
};
//判断ie9及其以下版本
var _IE = (function(){var v = 3, div = document.createElement('div'), all = div.getElementsByTagName('i');while (div.innerHTML = '<!--[if gt IE ' + (++v) + ']><i></i><![endif]-->',all[0]);return v > 4 ? v : false ;}());
//ajax判断是否登陆
function ajaxHasSession(yesCallback,noCallback){$.ajax({url:"/post/hassession",type:'POST',timeout:15000,success:function(data,textStatus){if(data==true){yesCallback();}else{noCallback();}}});}
//url参数拼接
function appendUrl(url,params,hook){if(params.length!=0){url+="?";for(key in params){url+=key+"="+params[key]+"&"}url=url.substring(0,url.length-1)}if(hook){url+="#"+hook}return url}
//根据windowsize判断标准还是宽屏
function setWindowSize(){if($.cookie("windowsize")==undefined||$.cookie("windowsize")=="null"||$.cookie("windowsize")==1){$("#j-window-df").attr("checked","checked");$("head").append("<style>.f-ct{width:80%;}</style>");}else{$("#j-window-lg").attr("checked","checked");$("head").append("<style>.f-ct{width:90%;}</style>");}};
//判断某个cookie是否为空
function cookieExist(key){
	if ($.cookie(key) == undefined || $.cookie(key) == "" || $.cookie("WOKUID") == "null" || $.cookie(key) == null){
		return false;
	}else{
		return true;
	}
}
//根据消息类别返回显示信息
function getMessageType(title){
	var result;
	switch (title){
		case "system": //系统消息
			result = "<span class='label label-primary'><i class='fa fa-flag'></i>&nbsp;系统信息</span>";
			break;
		case "game": //专区消息
			result = "<span class='label label-warning'><i class='fa fa-gamepad'></i>&nbsp;专区信息</span>";
			break;
	}
	return result;
}
//根据消息类别返回标题内容
function getMessageDescription(description,read,link){
	var result = description;
	if (read == false){
		result = "<span class='f-fwb'>"+description+"</span>";
	}
	if (link != ""){
		result += "&nbsp;&nbsp;<a class='f-fs14' href='"+link+"' target='_blank'>点击查看</a>";
	}
	return result;
}

/* jquery插件拓展 */
//bootstrap button拓展
+function($){var Button=function(element,options){this.$element=$(element);this.options=$.extend({},Button.DEFAULTS,options);this.isLoading=false};Button.VERSION="3.2.0";Button.DEFAULTS={loadingText:"loading..."};Button.prototype.setState=function(state){var d="disabled";var $el=this.$element;var val=$el.is("input")?"val":"html";var data=$el.data();state=state+"Text";if(data.resetText==null){$el.data("resetText",$el[val]())}$el[val](data[state]==null?this.options[state]:data[state]);setTimeout($.proxy(function(){if(state=="loadingText"){this.isLoading=true;$el.addClass(d).attr(d,d)}else{if(this.isLoading){this.isLoading=false;$el.removeClass(d).removeAttr(d)}}},this),0)};Button.prototype.toggle=function(){var changed=true;var $parent=this.$element.closest('[data-toggle="buttons"]');if($parent.length){var $input=this.$element.find("input");if($input.prop("type")=="radio"){if($input.prop("checked")&&this.$element.hasClass("active")){changed=false}else{$parent.find(".active").removeClass("active")}}if(changed){$input.prop("checked",!this.$element.hasClass("active")).trigger("change")}}if(changed){this.$element.toggleClass("active")}};function Plugin(option){return this.each(function(){var $this=$(this);var data=$this.data("bs.button");var options=typeof option=="object"&&option;if(!data){$this.data("bs.button",(data=new Button(this,options)))}if(option=="toggle"){data.toggle()}else{if(option){data.setState(option)}}})}var old=$.fn.button;$.fn.button=Plugin;$.fn.button.Constructor=Button;$.fn.button.noConflict=function(){$.fn.button=old;return this};function getBtnTarget(target){var $target=$(target);return $target.hasClass("btn")?$target:$target.parent(".btn")}$(document).on("click.bs.button.data-api",'[data-toggle^="button"]',function(e){var $btn=$(e.target);if(!$btn.hasClass("btn")){$btn=$btn.closest(".btn")}Plugin.call($btn,"toggle");e.preventDefault()}).on("focus.bs.button.data-api",'[data-toggle^="button"]',function(e){getBtnTarget(e.target).addClass("focus")}).on("blur.bs.button.data-api",'[data-toggle^="button"]',function(e){getBtnTarget(e.target).removeClass("focus")})}(jQuery);
//jquery ajax拓展
//var ajax=$.ajax;$.extend({ajax:function(url,options){typeof url=="object"&&(options=url,url=undefined),options=options||{},url=options.url;var xsrf,xsrflist;xsrf=$.cookie("_xsrf"),xsrflist=xsrf.split("|");var xsrftoken=base64_decode(xsrflist[0]),headers=options.headers||{},domain=document.domain.replace(/\./ig,"\\.");if(!/^(http:|https:).*/.test(url)||eval("/^(http:|https:)\\/\\/(.+\\.)*"+domain+".*/").test(url))headers=$.extend(headers,{"X-Xsrftoken":xsrftoken});return options.headers=headers,ajax(url,options)}})
//jquery date拓展 生成yyyymmdd形式日期
Date.prototype.yyyymmdd = function() {var yyyy = this.getFullYear().toString();var mm = (this.getMonth()+1).toString();var dd  = this.getDate().toString();return yyyy + (mm[1]?mm:"0"+mm[0]) + (dd[1]?dd:"0"+dd[0]);};
//生成yyyymm形式日期
Date.prototype.yyyymm = function() {var yyyy = this.getFullYear().toString();var mm = (this.getMonth()+1).toString();return yyyy + (mm[1]?mm:"0"+mm[0]);};
//placeholder ie下jquery拓展
$.fn.placeholder=function(){
	var i = document.createElement("input"),
	placeholdersupport = "placeholder" in i;
	if (!placeholdersupport) {
		$('[placeholder]').focus(function() {
        var input = $(this);
        if (input.val() == input.attr('placeholder')) {
            input.val('');
            input.removeClass('placeholder');
        }
    }).blur(function() {
        var input = $(this);
        if (input.val() == '' || input.val() == input.attr('placeholder')) {
            input.addClass('placeholder');
            input.val(input.attr('placeholder'));
        }
    }).blur();
	}
	return this
};
//自动判断相对绝对路径显示用户头像
$.fn.showUserImage=function(){return this.each(function(){var url=$(this).attr("userimage");if(url!=undefined&&url!=""){if(!isNaN(url)){$(this).attr("src","/static/img/user/"+url+".jpg")}else{$(this).attr("src",url)}}})};
//生成分页拓展
var Pagination = {};
(function($) {
	Pagination = function(element,options,callback){
		//初始化
		var defaults = {
			category: "blog",
			page: 1,
			allpage: 1,
			url: "http://www.wokugame.com",
			data: {},
			start: false //是否初始就发送ajax请求
		};
		var opts = $.extend(defaults, options);
		var loading = false; //是否在ajax中
		var _this = this;
		//几个重要方法
		this.overwrite = function(params){ //重写参数
			opts = $.extend(opts, params);
		}
		
		this.doAjax = function(button, e){ //发送ajax请求并调用回调方法
			if (button != null && e && button.attr("page") == 0) {
				return false
			}
			if (loading == true) {
				return false
			}
			var real_data = $.extend(opts.data, {page:opts.page});
			$.ajax({
				url: opts.url,
				type: "POST",
				data: real_data,
				beforeSend: function() {
					if (button != null) {
						if (e) {
							button.parent().addClass("disabled");
							button.button("loading")
						} else {
							button.attr("disabled", "disabled")
						}
					}
					loading = true
				},
				success: function(data, textStatus) {
					if (button != null) {
						if (e) {
							opts.page = button.attr("page")
						} else {
							opts.page = button.val()
						}
					}
					callback(data, opts.page)
				},
				complete: function() {
					if (button != null) {
						if (e) {
							button.button("reset");
							button.parent().removeClass("disabled")
						} else {
							button.removeAttr("disabled")
						}
					}
					loading = false
				}
			})
		}

		this.create = function(){ //创建分页DOM内容
			element.html("");
			if (opts.allpage == 1) {
				return;
			}
			opts.page = parseInt(opts.page);
			element.html("<ul class='g-pa'></ul>").css("position", "relative");
			var pagin = element.find(".g-pa");
			if (opts.page == 1) {
				pagin.append("<li><a class='disabled f-bln' href='javascript:void(0);' page='0'><i class='fa fa-arrow-left'></i></a></li>")
			} else {
				pagin.append("<li><a class='f-bln' href='javascript:void(0);' page='" + (opts.page - 1) + "'><i class='fa fa-arrow-left'></i></a></li>")
			}
			if (opts.allpage < 7) {
				for (var i = 1; i <= opts.allpage; i++) {
					if (i == opts.page) {
						pagin.append("<li><a class='active' href='javascript:void(0);' page='0'>" + i + "</a></li>")
					} else {
						pagin.append("<li><a href='javascript:void(0);' page='" + i + "'>" + i + "</a></li>")
					}
				}
			} else {
				if (opts.page < 6) {
					for (var i = 1; i <= 6; i++) {
						if (i == opts.page) {
							pagin.append("<li><a class='active' href='javascript:void(0);' page='0'>" + i + "</a></li>")
						} else {
							pagin.append("<li><a href='javascript:void(0);' page='" + i + "'>" + i + "</a></li>")
						}
					}
					pagin.append("<li><a class='disabled' href='javascript:void(0);' page='0'>...</a></li>");
					pagin.append("<li><a href='javascript:void(0);' page='" + opts.allpage + "'>" + opts.allpage + "</a></li>")
				} else {
					pagin.append("<li><a href='javascript:void(0);' page='1'>1</a></li>");
					pagin.append("<li><a href='javascript:void(0);' page='2'>2</a></li>");
					pagin.append("<li><a class='disabled' href='javascript:void(0);' page='0'>...</a></li>");
					if (opts.allpage - opts.page < 3) {
						for (var i = opts.allpage - 3; i <= opts.allpage; i++) {
							if (i == opts.page) {
								pagin.append("<li><a class='active' href='javascript:void(0);' page='0'>" + i + "</a></li>")
							} else {
								pagin.append("<li><a href='javascript:void(0);' page='" + i + "'>" + i + "</a></li>")
							}
						}
					} else {
						for (var i = opts.page - 2; i <= opts.page + 3; i++) {
							if (i == opts.page) {
								pagin.append("<li><a class='active' href='javascript:void(0);' page='0'>" + i + "</a></li>")
							} else {
								pagin.append("<li><a href='javascript:void(0);' page='" + i + "'>" + i + "</a></li>")
							}
						}
						pagin.append("<li><a class='disabled' href='javascript:void(0);' page='0'>...</a></li>");
						pagin.append("<li><a href='javascript:void(0);' page='" + (opts.allpage - 1) + "'>" + (opts.allpage - 1) + "</a></li>");
						pagin.append("<li><a href='javascript:void(0);' page='" + opts.allpage + "'>" + opts.allpage + "</a></li>")
					}
				}
			}
			if (opts.page == opts.allpage) {
				pagin.append("<li><a class='disabled f-brn' href='javascript:void(0);' page='0'><i class='fa fa-arrow-right'></i></a></li>")
			} else {
				pagin.append("<li><a class='f-brn' href='javascript:void(0);' page='" + (opts.page + 1) + "'><i class='fa fa-arrow-right'></i></a></li>")
			}
			element.append("<select class='g-pas'></select>");
			var select = element.find("select");
			for (var i = 1; i <= opts.allpage; i++) {
				if (i == opts.page) {
					select.append("<option value='" + i + "' selected>" + i + "&nbsp;页</option>")
				} else {
					select.append("<option value='" + i + "'>" + i + "&nbsp;页</option>")
				}
			}
			element.show();
		}
		
		this.refresh = function(){
			this.doAjax(null,callback);
		}
		
		//监听事件
		element.on("change", "select", function() {
			opts.page = parseInt($(this).val());
			_this.doAjax($(this), false);
		});
		element.on("click", "[page]", function() {
			opts.page = parseInt($(this).attr("page"));
			_this.doAjax($(this), true);
		});
		//是否启动初始调用
		if (opts.start) {
			_this.doAjax(null, true);
		}
		//返回自身对象
		return this;
	}
})(jQuery);

//倒计时拓展
var Timediff = {};
(function($) {
	Timediff = function(element,options,callback){
		//初始化
		var defaults = {
			second : 0
		};
		var opts = $.extend(defaults, options);
		
		this.run = function(){
			function Run(){
				var day=0,
		        hour=0,
		        minute=0,
		        second=0;//时间默认值        
		    if(defaults.second > 0){
		        day = Math.floor(defaults.second / (60 * 60 * 24));
		        hour = Math.floor(defaults.second / (60 * 60)) - (day * 24);
		        minute = Math.floor(defaults.second / 60) - (day * 24 * 60) - (hour * 60);
		        second = Math.floor(defaults.second) - (day * 24 * 60 * 60) - (hour * 60 * 60) - (minute * 60);
		    }else if (defaults.second == 0){
					callback();
				}
		    if (minute <= 9) minute = '0' + minute;
		    if (second <= 9) second = '0' + second;
		    element.find("#j-day").html(day+" 天");
		    element.find("#j-hour").html(hour+" 时");
		    element.find("#j-minute").html(minute+" 分");
		   	element.find("#j-second").html(second+" 秒");
		    defaults.second--;
			}
			var inter = setInterval(function(){
				if (!$.contains(document.body,element[0])){
					clearInterval(inter);
				}
		    Run();
				console.log("执行循环");
		    }, 1000);
			Run();
		}
	}
})(jQuery);

/* 通用执行 */
//初始化timeago组件
/*$.timeago.settings.strings = {
	prefixAgo: null,
  prefixFromNow: null,
  suffixAgo: "前",
  suffixFromNow: "现在开始",
  inPast: '现在',
  seconds: "1分钟",
  minute: "大约1分钟",
  minutes: "%d 分钟",
  hour: "大约1小时",
  hours: "%d 小时",
  day: "一天",
  days: "%d 天",
  month: "一个月",
  months: "%d 个月",
  year: "一年",
  years: "%d 年",
  wordSeparator: " ",
  numbers: []
}*/
//百度站长统计
var _bdhmProtocol = (("https:" == document.location.protocol) ? " https://" : " http://");
document.write(unescape("%3Cscript src='" + _bdhmProtocol + "hm.baidu.com/h.js%3F883b03c4cfad60bccf21ced9b32e8676' type='text/javascript'%3E%3C/script%3E"));
//防止被iframe
if (top.location != self.location)top.location=self.location;
//主窗体最小高度设定
var height = window.document.body.clientHeight - 200;
$("#j-main-content").css("min-height",height+"px");
//模拟placeholder
$("input").placeholder();
//刷新菜单信息
refreshNav();
//每过5分钟请求一次是否有新消息
setInterval(function(){
	refreshNav();
},300000);

setTimeout(function(){
	refreshNav();
},1000);

/* 导航条处理 */
//鼠标移动到drop出现下拉菜单
var header_message = false;
$(".j-drop").hover(function(){
	$(this).find(".j-drop-content").show();
	var _this = $(this);
	switch($(this).attr("special")){
		case "message": //消息，当未读消息不为0，或者消息栏为空时，获取消息
			if ($.cookie("WOKUREAD") > 0 || header_message == false){
				var newNumber = parseInt($.cookie("WOKUREAD"));
				header_message = true;
				$.cookie("WOKUREAD",0,{expires: 3650, path: '/'});
				$(".m-nav .info-number").text("").hide();
				$.ajax({
					url: "/account_messagepost",
					type: "POST",
					data: {page:1,clear:"true"},
					beforeSend:function(){
						_this.find(".j-drop-content").html("<li class='f-tac f-p20 text-muted'>消息获取中&nbsp;<i class='fa fa-refresh fa-spin'></i></li>");
					},
					success: function(data, textStatus) {
						var c = _this.find(".j-drop-content");
						c.html("");
						if (data == ""){
							c.html("<li class='f-tac f-p20 text-muted'>暂无消息</li>");
							return false;
						}
						for (var i=0;i<data.length;i++){
							var title = getMessageType(data[i].Message.Type,data[i].Message.Link);
							var description = "<span class='f-fwb f-ml10'>"+data[i].Message.Description+"</span><span class='f-ml10 timeago' title='"+data[i].Time+"'></span>";
							var link = "<a href='"+data[i].Message.Link+"' target='_blank'>点击查看</a>";
							c.append("<div id='message-"+i+"' class='message-content g-bd1 f-cb'><a href='/user.html?to=account_message' class='g-sd1' target='_blank'>"+title+description+"</a><div class='g-mn1'><div class='g-mn1c f-fs14'>"+link+"</div></div></div>");
							if (i != data.length-1){
								c.append("<li class='cut'></li>");
							}
							//显示友好时间
							$(".timeago").timeago();
						}
						//让最新的消息闪一下
						for (var i=0;i<newNumber;i++){
							$("#message-"+i).css("background","black");
						}
						setTimeout(function(){
							$(".m-nav .message-content").removeAttr("style");
						},500);
					}
				});
			}
			break;
	}
},function(){
	$(this).find(".j-drop-content").hide();
});
//退出按钮点击
$(".m-nav .j-logined #j-out").click(function(){
	$.ajax({
		url: "/exist",
		type: "POST",
		timeout: 15000,
		success: function(){
			//显示登陆注册按钮
			$(".m-nav #j-login").show()//隐藏登陆按钮
			$(".m-nav #j-register").show()//隐藏注册按钮
			//隐藏已登录按钮
			$(".m-nav .j-logined").hide();
			//删除本地cookie
			$.cookie("WOKUID", null, {path: '/'});
		}
	});
});
//实例化date
var date = new Date();
//如果是新的一天则重置数据
if (date.yyyymmdd() > $.cookie("yyyymmdd") || $.cookie("yyyymmdd") == "null"){ //是新的一天
	$.cookie("WOKUUPSIZE",0,{expires: 3650, path: '/'});//每日上传量清零
}
//如果是新的一月则重置数据
if (date.yyyymm() > $.cookie("yyyymm") || $.cookie("yyyymm") == "null"){ //是新的一月
	$.cookie("WOKUFREE",0,{expires: 3650, path: '/'});//每月免费额度清零
}
//刷新日期
$.cookie("yyyymmdd",date.yyyymmdd(),{expires: 3650, path: '/'});
$.cookie("yyyymm",date.yyyymm(),{expires: 3650, path: '/'});
//头部修改屏幕按钮事件
setWindowSize();
$("input[name=windowsize]").change(function(){
	switch($("input[name=windowsize]:checked").attr("id")){
	  case "j-window-df":
			$.cookie("windowsize",1,{expires: 3650, path: '/'});
			break;
	  case "j-window-lg":
			$.cookie("windowsize",2,{expires: 3650, path: '/'});
			break;
	  default:
			break;
	}
	setWindowSize();
});

/* 页尾处理 */
//鼠标移动到二维码上显示二维码图标
$(".g-ft #j-qrcode").hover(function(){
	$(this).find("img").show();
},function(){
	$(this).find("img").hide();
});