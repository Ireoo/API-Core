(function(e){function t(t){for(var o,r,i=t[0],l=t[1],c=t[2],d=0,f=[];d<i.length;d++)r=i[d],Object.prototype.hasOwnProperty.call(s,r)&&s[r]&&f.push(s[r][0]),s[r]=0;for(o in l)Object.prototype.hasOwnProperty.call(l,o)&&(e[o]=l[o]);u&&u(t);while(f.length)f.shift()();return a.push.apply(a,c||[]),n()}function n(){for(var e,t=0;t<a.length;t++){for(var n=a[t],o=!0,i=1;i<n.length;i++){var l=n[i];0!==s[l]&&(o=!1)}o&&(a.splice(t--,1),e=r(r.s=n[0]))}return e}var o={},s={app:0},a=[];function r(t){if(o[t])return o[t].exports;var n=o[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,r),n.l=!0,n.exports}r.m=e,r.c=o,r.d=function(e,t,n){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},r.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)r.d(n,o,function(t){return e[t]}.bind(null,o));return n},r.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],l=i.push.bind(i);i.push=t,i=i.slice();for(var c=0;c<i.length;c++)t(i[c]);var u=l;a.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"0009":function(e,t,n){"use strict";n.r(t);var o=n("2b0e"),s=function(){};o["default"].prototype.console=s,t["default"]=s},"02bb":function(e,t,n){},"034f":function(e,t,n){"use strict";n("64a9")},"0392":function(e,t,n){},1:function(e,t){},10:function(e,t){},11:function(e,t){},1179:function(e,t,n){"use strict";n.r(t);var o=n("2b0e"),s=n("c64e");o["default"].prototype.$uuid=function(){return s()},t["default"]=o["default"]},12:function(e,t){},"12ba":function(e,t,n){"use strict";n("c472")},13:function(e,t){},"13ac":function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col",[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],staticStyle:{width:"100%"},attrs:{data:e.infos,border:""}},[n("el-table-column",{attrs:{prop:"content",label:"内容"}}),n("el-table-column",{attrs:{prop:"from",label:"来源"}}),n("el-table-column",{attrs:{label:"创建时间"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("i",{staticClass:"el-icon-time"}),n("span",{staticStyle:{"margin-left":"10px"}},[e._v(e._s(e.$time(t.row.createTime).format("YYYY-MM-DD")))])]}}])}),n("el-table-column",{attrs:{fixed:"right",label:"操作",width:"200"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{attrs:{loading:t.row.update,disabled:t.row.del,type:"primary",size:"mini"},on:{click:function(n){return e.secret(t.row)}}},[e._v("重置密钥")]),n("el-button",{attrs:{loading:t.row.del,disabled:t.row.update,type:"danger",size:"mini"},on:{click:function(n){return e.del(t)}}},[e._v("删除")])]}}])})],1)],1)],1)},s=[],a={name:"info",title:"通知",components:{},data:function(){return{infos:[],loading:!1}},mounted:function(){}},r=a,i=n("2877"),l=Object(i["a"])(r,o,s,!1,null,"0257e85f",null);t["default"]=l.exports},14:function(e,t){},1441:function(e,t,n){"use strict";n("0392")},"185b":function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col")],1)},s=[],a={name:"password",title:"修改密码",icon:"iconfont icon-password",index:0},r=a,i=n("2877"),l=Object(i["a"])(r,o,s,!1,null,"9cfde1f8",null);t["default"]=l.exports},"1b03":function(e,t,n){"use strict";n.r(t);var o=n("2b0e"),s=n("5c96"),a=n.n(s);n("0fae");o["default"].use(a.a),t["default"]=o["default"]},"1cee":function(e,t,n){"use strict";n.r(t);var o=n("2b0e");o["default"].prototype.DEBUG=!1,t["default"]=o["default"]},2:function(e,t){},"288b":function(e,t,n){},2948:function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",{directives:[{name:"loading",rawName:"v-loading",value:e.fullscreenLoading,expression:"fullscreenLoading"}]},[n("el-col",[n("h1",[e._v(e._s(e.app.title))]),n("div",{staticClass:"desc"},[e._v("\n      创建于: "+e._s(e.$time(e.app.createTime).format("YYYY-MM-DD"))+" 最后更新:\n      "+e._s(e.$time(e.app.updateTime).format("YYYY-MM-DD"))+"\n    ")]),n("div",{staticClass:"desc"},[e._v("\n      密钥:\n      "),n("b",[e._v(e._s(e.app.secret))])])]),n("el-col",[n("el-tabs",{attrs:{"tab-position":"top"}},[n("el-tab-pane",{attrs:{label:"数据表"}},[0===e.collections.data.length?n("div",{staticClass:"desc"},[e._v("\n          暂时没有添加任务数据表,请使用接口创建自己的数据吧!\n        ")]):e._e(),0!==e.collections.data.length?n("ul",e._l(e.collections.data,(function(t){return n("li",{key:t.info.uuid},[e._v("\n            "+e._s(t.name)+"\n          ")])})),0):e._e()]),n("el-tab-pane",{attrs:{label:"监控"}})],1)],1)],1)},s=[],a={name:"id",title:"应用信息",props:["collapse"],data:function(){return{app:{},collections:{loading:!1,data:[]},fullscreenLoading:!1}},created:function(){var e=this;this.fullscreenLoading=!0,this.console(this.$route.params.id),this.$http("apps/once",{where:{_id:this.$route.params.id,uuid:this.$store.state.user.info._id},other:{show:{_id:0,uuid:0}}}).then((function(t){t?(e.app=t,document.title=e.app.title):(e.$message.error("应用不存在!"),e.$router.push({path:"/app"})),e.fullscreenLoading=!1})).catch((function(t){e.$message.error(t)})),this.collections={loading:!0,data:[]},this.$http("all/listCollections",{app:this.$route.params.id}).then((function(t){e.console(t),e.collections={loading:!1,data:t}}))}},r=a,i=(n("3afa"),n("2877")),l=Object(i["a"])(r,o,s,!1,null,"11537659",null);t["default"]=l.exports},"2fdb":function(e,t,n){},3:function(e,t){},"326c":function(e,t,n){var o={"./ElementUI":"1b03","./ElementUI.js":"1b03","./axios":"be3b","./axios.js":"be3b","./console":"0009","./console.js":"0009","./debug":"1cee","./debug.js":"1cee","./md5":"be72","./md5.js":"be72","./moment":"894f","./moment.js":"894f","./uuid":"1179","./uuid.js":"1179"};function s(e){var t=a(e);return n(t)}function a(e){if(!n.o(o,e)){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t}return o[e]}s.keys=function(){return Object.keys(o)},s.resolve=a,e.exports=s,s.id="326c"},"3afa":function(e,t,n){"use strict";n("2fdb")},4:function(e,t){},"41ab":function(e,t,n){"use strict";n("5248")},"45b4":function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col",{staticStyle:{"text-align":"right","margin-bottom":"20px"}},[n("el-button-group",[n("el-button",{attrs:{type:e.list?"":"primary",size:"mini",icon:"iconfont icon-list1"},on:{click:function(t){e.list=!1}}}),n("el-button",{attrs:{type:e.list?"primary":"",size:"mini",icon:"iconfont icon-list"},on:{click:function(t){e.list=!0}}})],1)],1),e.list?n("el-col",[n("el-row",[n("el-col",[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],staticStyle:{width:"100%"},attrs:{data:e.appsShow,border:""}},[n("el-table-column",{attrs:{prop:"title",label:"名称"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("router-link",{attrs:{to:"/app/info/"+t.row._id}},[e._v(e._s(t.row.title))])]}}],null,!1,4212912267)}),n("el-table-column",{attrs:{prop:"_id",label:"APP id"}}),n("el-table-column",{attrs:{prop:"secret",label:"APP secret"}}),n("el-table-column",{attrs:{label:"创建时间"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("i",{staticClass:"el-icon-time"}),n("span",{staticStyle:{"margin-left":"10px"}},[e._v(e._s(e.$time(t.row.createTime).format("YYYY-MM-DD")))])]}}],null,!1,2536342592)}),n("el-table-column",{attrs:{fixed:"right",label:"操作",width:"80"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{attrs:{loading:t.row.update,disabled:t.row.del,type:"primary",size:"mini"},on:{click:function(n){return e.secret(t.row)}}},[e._v("编辑")])]}}],null,!1,1320399595)})],1)],1)],1)],1):e._e(),e.list?e._e():n("el-col",[n("el-row",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],attrs:{gutter:20}},e._l(e.appsShow,(function(t){return n("el-col",{key:t._id,staticStyle:{"margin-bottom":"20px"},attrs:{span:e.collapse?24:12}},[n("el-card",{attrs:{shadow:"hover"}},[n("div",{staticClass:"clearfix",staticStyle:{position:"relative"},attrs:{slot:"header"},slot:"header"},[n("span",[n("router-link",{attrs:{to:"/app/info/"+t._id}},[e._v(e._s(t.title))])],1),n("span",{staticClass:"desc m5"},[e._v("( ID: "+e._s(t._id)+" )")]),n("el-button",{staticStyle:{position:"absolute",top:"-5px",right:"0"},attrs:{type:"",size:"mini",loading:t.update,disabled:t.del},on:{click:function(n){return e.secret(t)}}},[e._v("编辑")])],1),n("div",{staticClass:"desc item"},[e._v(e._s(t.desc))]),n("div",{staticClass:"text item"},[e._v("密钥: "+e._s(t.secret))]),n("div",{staticClass:"text item"},[e._v("创建于: "+e._s(e.$time(t.createTime).format("YYYY-MM-DD")))])])],1)})),1)],1)],1)},s=[],a={name:"admin",title:"应用管理",icon:"iconfont icon-gaikuang",index:0,show:!0,props:["collapse"],components:{},data:function(){return{apps:[],loading:!0,list:!1}},methods:{secret:function(e){this.console(e),this.$router.push({path:"/app/edit/".concat(e._id)})},del:function(e){var t=this;this.console(e),e.del=!0,this.$http("apps/remove",{where:{_id:e._id}}).then((function(){t.$message.success("删除成功!"),e.show=!1,e.del=!1})).catch((function(e){t.$message.error(e)}))}},computed:{appsShow:function(){return this.apps.filter((function(e){return e.show}))}},mounted:function(){var e=this;this.$http("apps/find",{where:{uuid:this.$store.state.user.info._id}}).then((function(t){e.loading=!1,e.apps=t.map((function(e){return e.del=!1,e.update=!1,e.show=!0,e}))})).catch((function(t){e.loading=!1,e.$message.error(t)}))}},r=a,i=(n("d84b"),n("2877")),l=Object(i["a"])(r,o,s,!1,null,"156d0758",null);t["default"]=l.exports},4649:function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",{staticClass:"main",attrs:{type:"flex",justify:"center"}},[n("el-col",{staticStyle:{"min-width":"300px","max-width":"400px"},attrs:{span:12}},[n("el-row",{staticClass:"form"},[n("el-col",[n("label",{class:{focus:e.focus.username}},[n("span",[e._v("用户名")]),n("el-input",{ref:"username",attrs:{placeholder:e.focus.username?e.placeholder.username:"",disabled:e.loading},on:{focus:function(t){e.focus.username=!0},blur:function(t){e.focus.username=""!==e.form.username}},model:{value:e.form.username,callback:function(t){e.$set(e.form,"username",t)},expression:"form.username"}})],1)]),n("el-col",[n("label",{class:{focus:e.focus.password}},[n("span",[e._v("密码")]),n("el-input",{ref:"password",attrs:{"show-password":"",placeholder:e.focus.password?e.placeholder.password:"",disabled:e.loading},on:{focus:function(t){e.focus.password=!0},blur:function(t){e.focus.password=""!==e.form.password}},model:{value:e.form.password,callback:function(t){e.$set(e.form,"password",t)},expression:"form.password"}})],1)]),n("el-col",[n("label",{class:{focus:e.focus.passwordTwo}},[n("span",[e._v("确认密码")]),n("el-input",{ref:"passwordTwo",attrs:{"show-password":"",placeholder:e.focus.passwordTwo?e.placeholder.passwordTwo:"",disabled:e.loading},on:{focus:function(t){e.focus.passwordTwo=!0},blur:function(t){e.focus.passwordTwo=""!==e.form.passwordTwo}},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.login.apply(null,arguments)}},model:{value:e.form.passwordTwo,callback:function(t){e.$set(e.form,"passwordTwo",t)},expression:"form.passwordTwo"}})],1)]),n("el-col",[n("el-row",{attrs:{type:"flex",align:"bottom"}},[n("el-col",{attrs:{span:12}},[n("el-button",{attrs:{type:"primary",loading:e.loading},on:{click:e.reg}},[e._v("注册")])],1),n("el-col",{staticClass:"text-right",attrs:{span:12}},[n("router-link",{attrs:{to:"/login"}},[e._v("登陆")])],1)],1)],1)],1)],1)],1)},s=[],a={name:"reg",title:"注册帐号",data:function(){return{focus:{username:!1,password:!1,passwordTwo:!1},placeholder:{username:"手机号/邮箱",password:"格式:[A-Z]{1}(.*){7,}",passwordTwo:"再次输入一遍"},form:{username:"",password:"",passwordTwo:""},loading:!1}},methods:{reg:function(){var e=this;if(""===this.form.username)return this.$refs.username.focus();if(""===this.form.password)return this.$refs.password.focus();if(""===this.form.passwordTwo)return this.$refs.passwordTwo.focus();this.loading=!0;var t=JSON.parse(JSON.stringify(this.form));t.password===t.passwordTwo&&(t.password=this.$md5(t.password),delete t.passwordTwo,this.$http("users/insert",{data:t}).then((function(n){e.$message.success("".concat(t.username," 用户注册成功!")),e.loading=!1,e.$store.commit("account.UPDATE",n)})).catch((function(t){switch(t.code){case 11e3:e.$message.error("该用户已经存在,请尝试更换用户名!");break;default:e.$message.error("未知错误: ".concat(JSON.stringify(t)));break}e.loading=!1})))}},mounted:function(){}},r=a,i=(n("d2d8"),n("2877")),l=Object(i["a"])(r,o,s,!1,null,"1ef755eb",null);t["default"]=l.exports},4678:function(e,t,n){var o={"./af":"2bfb","./af.js":"2bfb","./ar":"8e73","./ar-dz":"a356","./ar-dz.js":"a356","./ar-kw":"423e","./ar-kw.js":"423e","./ar-ly":"1cfd","./ar-ly.js":"1cfd","./ar-ma":"0a84","./ar-ma.js":"0a84","./ar-sa":"8230","./ar-sa.js":"8230","./ar-tn":"6d83","./ar-tn.js":"6d83","./ar.js":"8e73","./az":"485c","./az.js":"485c","./be":"1fc1","./be.js":"1fc1","./bg":"84aa","./bg.js":"84aa","./bm":"a7fa","./bm.js":"a7fa","./bn":"9043","./bn-bd":"9686","./bn-bd.js":"9686","./bn.js":"9043","./bo":"d26a","./bo.js":"d26a","./br":"6887","./br.js":"6887","./bs":"2554","./bs.js":"2554","./ca":"d716","./ca.js":"d716","./cs":"3c0d","./cs.js":"3c0d","./cv":"03ec","./cv.js":"03ec","./cy":"9797","./cy.js":"9797","./da":"0f14","./da.js":"0f14","./de":"b469","./de-at":"b3eb","./de-at.js":"b3eb","./de-ch":"bb71","./de-ch.js":"bb71","./de.js":"b469","./dv":"598a","./dv.js":"598a","./el":"8d47","./el.js":"8d47","./en-au":"0e6b","./en-au.js":"0e6b","./en-ca":"3886","./en-ca.js":"3886","./en-gb":"39a6","./en-gb.js":"39a6","./en-ie":"e1d3","./en-ie.js":"e1d3","./en-il":"7333","./en-il.js":"7333","./en-in":"ec2e","./en-in.js":"ec2e","./en-nz":"6f50","./en-nz.js":"6f50","./en-sg":"b7e9","./en-sg.js":"b7e9","./eo":"65db","./eo.js":"65db","./es":"898b","./es-do":"0a3c","./es-do.js":"0a3c","./es-mx":"b5b7","./es-mx.js":"b5b7","./es-us":"55c9","./es-us.js":"55c9","./es.js":"898b","./et":"ec18","./et.js":"ec18","./eu":"0ff2","./eu.js":"0ff2","./fa":"8df4","./fa.js":"8df4","./fi":"81e9","./fi.js":"81e9","./fil":"d69a","./fil.js":"d69a","./fo":"0721","./fo.js":"0721","./fr":"9f26","./fr-ca":"d9f8","./fr-ca.js":"d9f8","./fr-ch":"0e49","./fr-ch.js":"0e49","./fr.js":"9f26","./fy":"7118","./fy.js":"7118","./ga":"5120","./ga.js":"5120","./gd":"f6b4","./gd.js":"f6b4","./gl":"8840","./gl.js":"8840","./gom-deva":"aaf2","./gom-deva.js":"aaf2","./gom-latn":"0caa","./gom-latn.js":"0caa","./gu":"e0c5","./gu.js":"e0c5","./he":"c7aa","./he.js":"c7aa","./hi":"dc4d","./hi.js":"dc4d","./hr":"4ba9","./hr.js":"4ba9","./hu":"5b14","./hu.js":"5b14","./hy-am":"d6b6","./hy-am.js":"d6b6","./id":"5038","./id.js":"5038","./is":"0558","./is.js":"0558","./it":"6e98","./it-ch":"6f12","./it-ch.js":"6f12","./it.js":"6e98","./ja":"079e","./ja.js":"079e","./jv":"b540","./jv.js":"b540","./ka":"201b","./ka.js":"201b","./kk":"6d79","./kk.js":"6d79","./km":"e81d","./km.js":"e81d","./kn":"3e92","./kn.js":"3e92","./ko":"22f8","./ko.js":"22f8","./ku":"2421","./ku.js":"2421","./ky":"9609","./ky.js":"9609","./lb":"440c","./lb.js":"440c","./lo":"b29d","./lo.js":"b29d","./lt":"26f9","./lt.js":"26f9","./lv":"b97c","./lv.js":"b97c","./me":"293c","./me.js":"293c","./mi":"688b","./mi.js":"688b","./mk":"6909","./mk.js":"6909","./ml":"02fb","./ml.js":"02fb","./mn":"958b","./mn.js":"958b","./mr":"39bd","./mr.js":"39bd","./ms":"ebe4","./ms-my":"6403","./ms-my.js":"6403","./ms.js":"ebe4","./mt":"1b45","./mt.js":"1b45","./my":"8689","./my.js":"8689","./nb":"6ce3","./nb.js":"6ce3","./ne":"3a39","./ne.js":"3a39","./nl":"facd","./nl-be":"db29","./nl-be.js":"db29","./nl.js":"facd","./nn":"b84c","./nn.js":"b84c","./oc-lnc":"167b","./oc-lnc.js":"167b","./pa-in":"f3ff","./pa-in.js":"f3ff","./pl":"8d57","./pl.js":"8d57","./pt":"f260","./pt-br":"d2d4","./pt-br.js":"d2d4","./pt.js":"f260","./ro":"972c","./ro.js":"972c","./ru":"957c","./ru.js":"957c","./sd":"6784","./sd.js":"6784","./se":"ffff","./se.js":"ffff","./si":"eda5","./si.js":"eda5","./sk":"7be6","./sk.js":"7be6","./sl":"8155","./sl.js":"8155","./sq":"c8f3","./sq.js":"c8f3","./sr":"cf1e","./sr-cyrl":"13e9","./sr-cyrl.js":"13e9","./sr.js":"cf1e","./ss":"52bd","./ss.js":"52bd","./sv":"5fbd","./sv.js":"5fbd","./sw":"74dc","./sw.js":"74dc","./ta":"3de5","./ta.js":"3de5","./te":"5cbb","./te.js":"5cbb","./tet":"576c","./tet.js":"576c","./tg":"3b1b","./tg.js":"3b1b","./th":"10e8","./th.js":"10e8","./tk":"5aff","./tk.js":"5aff","./tl-ph":"0f38","./tl-ph.js":"0f38","./tlh":"cf75","./tlh.js":"cf75","./tr":"0e81","./tr.js":"0e81","./tzl":"cf51","./tzl.js":"cf51","./tzm":"c109","./tzm-latn":"b53d","./tzm-latn.js":"b53d","./tzm.js":"c109","./ug-cn":"6117","./ug-cn.js":"6117","./uk":"ada2","./uk.js":"ada2","./ur":"5294","./ur.js":"5294","./uz":"2e8c","./uz-latn":"010e","./uz-latn.js":"010e","./uz.js":"2e8c","./vi":"2921","./vi.js":"2921","./x-pseudo":"fd7e","./x-pseudo.js":"fd7e","./yo":"7f33","./yo.js":"7f33","./zh-cn":"5c3a","./zh-cn.js":"5c3a","./zh-hk":"49ab","./zh-hk.js":"49ab","./zh-mo":"3a6c","./zh-mo.js":"3a6c","./zh-tw":"90ea","./zh-tw.js":"90ea"};function s(e){var t=a(e);return n(t)}function a(e){if(!n.o(o,e)){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t}return o[e]}s.keys=function(){return Object.keys(o)},s.resolve=a,e.exports=s,s.id="4678"},"4b3b":function(e,t,n){var o={"./account":"b5c4","./account.vue":"b5c4","./account/password":"185b","./account/password.vue":"185b","./app/add":"8bb8","./app/add.vue":"8bb8","./app/database":"be9b","./app/database.vue":"be9b","./app/edit/#id":"da44","./app/edit/#id.vue":"da44","./app/info/#id":"2948","./app/info/#id.vue":"2948","./apps":"45b4","./apps.vue":"45b4","./home":"6511","./home.vue":"6511","./info":"13ac","./info.vue":"13ac","./login":"dd7b","./login.vue":"dd7b","./reg":"4649","./reg.vue":"4649"};function s(e){var t=a(e);return n(t)}function a(e){if(!n.o(o,e)){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t}return o[e]}s.keys=function(){return Object.keys(o)},s.resolve=a,e.exports=s,s.id="4b3b"},5:function(e,t){},5248:function(e,t,n){},"56d7":function(e,t,n){"use strict";n.r(t);n("a481"),n("ac6a"),n("cadf"),n("551c"),n("f751"),n("097d");var o=n("2b0e"),s=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-container",{ref:"app",class:["wrapper",{ios:e.ios}],attrs:{id:"app"}},[n("el-main",[n("el-scrollbar",[n("Header",{attrs:{collapse:e.collapse}}),n("router-view",{class:["content",{collapse:e.collapse},{body:!e.$route.meta.show}],attrs:{collapse:e.collapse}}),n("el-footer",[n("Footer",{attrs:{width:e.width}})],1)],1)],1)],1)},a=[],r=(n("4917"),function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col",{class:{collapse:e.collapse}},[n("ul",{staticClass:"left"},[e.$store.state.user.info?e._e():n("li",[n("router-link",{attrs:{to:"/home"}},[n("i",{staticClass:"iconfont icon-shouquanjiekou logo"}),n("span",[e._v(e._s(e.collapse?"":"爱数据"))])])],1),e.$store.state.user.info?n("li",[n("router-link",{attrs:{to:"/apps"}},[n("i",{staticClass:"iconfont icon-shouquanjiekou logo"}),n("span",[e._v(e._s(e.collapse?"":"爱数据"))])])],1):e._e()]),e.$store.state.user.info?e._e():n("ul",{staticClass:"right no"},[n("li",[n("router-link",{attrs:{to:"/login"}},[n("span",[e._v("登陆")])])],1),n("li",[n("router-link",{attrs:{to:"/reg"}},[n("span",[e._v("注册")])])],1)]),e.$store.state.user.info?n("ul",{staticClass:"right"},[n("li",[n("router-link",{attrs:{to:"/info"}},[0===e.infos?n("i",{staticClass:"iconfont icon-icon"}):e._e(),e.infos>0?n("el-badge",{staticClass:"item",attrs:{"is-dot":""}},[n("i",{staticClass:"iconfont icon-icon"})]):e._e()],1)],1),n("li",[n("el-dropdown",{attrs:{trigger:"click"},on:{command:e.routerGo}},[n("span",{staticClass:"el-dropdown-link"},[n("i",{staticClass:"iconfont icon-yingyong1"}),n("i",{staticClass:"el-icon-arrow-down el-icon--right"})]),n("el-dropdown-menu",{attrs:{slot:"dropdown"},slot:"dropdown"},[n("el-dropdown-item",{attrs:{command:"/app/add"}},[e._v("创建一个应用")]),n("el-dropdown-item",{attrs:{command:"/apps"}},[e._v("管理所有应用")])],1)],1)],1),n("li",[n("el-dropdown",{attrs:{trigger:"click"},on:{command:e.routerGo}},[n("span",{staticClass:"el-dropdown-link"},[n("img",{attrs:{src:"https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"}}),n("i",{staticClass:"el-icon-arrow-down el-icon--right"})]),n("el-dropdown-menu",{attrs:{slot:"dropdown"},slot:"dropdown"},[n("el-dropdown-item",{attrs:{command:"/account"}},[e._v("信息")]),n("el-dropdown-item",{attrs:{command:"/account/setting"}},[e._v("设置")]),n("el-dropdown-item",{attrs:{command:"logout"}},[e._v("退出")])],1)],1)],1)]):e._e(),n("br")])],1)}),i=[],l=(n("3b2b"),n("7f7f"),n("55dd"),{name:"Header",props:["collapse"],methods:{routerGo:function(e){"logout"===e?(this.$store.commit("account.REMOVE"),location.href="/"):this.$router.push({path:e})}},data:function(){return{isCollapse:!1,show:["App"],infos:0}},computed:{routes:function(){var e=this,t=this.$router.options.routes.filter((function(t){return e.show.indexOf(t.name)>-1})).sort((function(e,t){return e.meta.index-t.meta.index})).map((function(t){var n=JSON.parse(JSON.stringify(t));return n.child=e.$router.options.routes.filter((function(e){return new RegExp("^".concat(n.path,"/")).test(e.path)})).sort((function(e,t){return e.meta.index-t.meta.index})),n}));return t}}}),c=l,u=(n("41ab"),n("2877")),d=Object(u["a"])(c,r,i,!1,null,"3f4cbdd6",null),f=d.exports,p=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col",[n("el-menu",{class:["sider",{collapse:e.collapse}],attrs:{"default-active":e.$route.path,"background-color":"#545c64","text-color":"#fff","active-text-color":"#ffd04b",router:!0,collapse:e.collapse,"menu-trigger":"click"}},e._l(e.routes,(function(t,o){return n("el-menu-item-group",{key:o,attrs:{index:t.path}},[n("template",{slot:"title"},[e._v(e._s(e.collapse?"":t.meta.title))]),n("el-menu-item",{attrs:{index:t.path}},[n("i",{class:t.meta.icon}),n("span",{attrs:{slot:"title"},slot:"title"},[e._v(e._s(t.meta.title))])]),e._l(t.child,(function(t,o){return n("el-menu-item",{key:o,attrs:{index:t.path}},[n("i",{class:t.meta.icon}),n("span",{attrs:{slot:"title"},slot:"title"},[e._v(e._s(t.meta.title))])])}))],2)})),1)],1)],1)},m=[],b={name:"Sider",props:["collapse"],data:function(){return{show:["App","Account"]}},methods:{logout:function(){this.$store.commit("account.REMOVE")}},computed:{routes:function(){var e=this,t=this.$router.options.routes.filter((function(t){return e.show.indexOf(t.name)>-1})).sort((function(e,t){return e.meta.index-t.meta.index})).map((function(t){var n=JSON.parse(JSON.stringify(t));return n.child=e.$router.options.routes.filter((function(e){return new RegExp("^".concat(n.path,"/")).test(e.path)})).sort((function(e,t){return e.meta.index-t.meta.index})),n}));return t}},mounted:function(){}},h=b,g=(n("6d12"),Object(u["a"])(h,p,m,!1,null,"d9b01882",null)),v=g.exports,w=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[e.DEBUG?n("el-col",[e._v(e._s(e.$route.meta))]):e._e(),n("el-col",[e._v("© "+e._s((new Date).getFullYear())+" qiyi.io")])],1)},j=[],_={name:"Footer",props:["width"],data:function(){return{}}},y=_,x=(n("1441"),Object(u["a"])(y,w,j,!1,null,"67ce97b1",null)),k=x.exports,$={name:"App",data:function(){return{uri:"",width:document.documentElement.clientWidth,mini:document.documentElement.clientWidth<900,collapse:document.documentElement.clientWidth<700,ios:!!navigator.userAgent.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/)}},components:{Header:f,Sider:v,Footer:k},watch:{"$route.meta":{handler:function(e){this.console(e),e.title&&(document.title=e.title)},deep:!0}},mounted:function(){var e=this;this.$route.meta.title&&(document.title=this.$route.meta.title),window.onresize=function(){e.collapse=document.documentElement.clientWidth<700,e.mini=document.documentElement.clientWidth<900,e.width=document.documentElement.clientWidth}}},O=$,E=(n("034f"),n("12ba"),Object(u["a"])(O,s,a,!1,null,"17e2c8ca",null)),S=E.exports,C=n("2f62");o["default"].use(C["a"]);var z=n("e75c"),T={};z.keys().forEach((function(e){T[e.replace(/(\.\/|\.js)/g,"")]=z(e).default}));var D=new C["a"].Store({modules:T,strict:!1}),N=(n("28a5"),n("8c4f"));o["default"].use(N["a"]);var M=[{name:"Index",path:"/",redirect:"/home",meta:{title:"首页重定向",file:"To /home"}}],U=n("c321");U.keys().forEach((function(e){var t=e.replace(/(\.\/|\.vue)/g,""),o=t.replace("#",":"),s=t.split("/"),a=s.pop().replace(/\b([a-zA-Z#])(\w*)/g,(function(e,t,n){return t.replace("#","").toUpperCase()+n.toLowerCase()})).replace("#",""),r=s.join("/"),i=s.map((function(e){return e.replace(/\b([a-zA-Z#])(\w*)/g,(function(e,t,n){return t.toUpperCase()+n.toLowerCase()}))})).join(""),l=n("4b3b")("./".concat(t)).default,c={path:"index"===a.toLowerCase()?"/".concat(r):"/".concat(o.toLowerCase()),name:"index"===a.toLowerCase()&&""!==i?i:a,component:l,meta:{title:l.title,index:0|l.index,icon:l.icon||""}};M.push(c)})),M.push({path:"*",redirect:"/"});var A={routes:M},Y=new N["a"](A);Y.beforeEach((function(e,t,n){/^\/[app|account]/.test(e.path)?(e.meta.show=!0,D.state.user.info?n():n({path:"/login",query:{redirect:e.fullPath}})):n()}));var J=n("f666");J.keys().forEach((function(e){n("326c")("./".concat(e.replace(/(\.\/|\.js)/g,"")))})),o["default"].config.productionTip=!1,new o["default"]({router:Y,store:D,render:function(e){return e(S)}}).$mount("#app")},6:function(e,t){},"64a9":function(e,t,n){},6511:function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col")],1)},s=[],a={name:"home",title:"爱数据"},r=a,i=n("2877"),l=Object(i["a"])(r,o,s,!1,null,"3a4573a1",null);t["default"]=l.exports},"6d12":function(e,t,n){"use strict";n("8988")},7:function(e,t){},"7f34":function(e,t,n){"use strict";n.r(t);var o,s=n("ade3"),a=n("e0ac"),r={info:null};localStorage.getItem("account")&&(r.info=JSON.parse(Object(a["decrypt"])(localStorage.getItem("account"),"admio")));var i=(o={},Object(s["a"])(o,"account.UPDATE",(function(e,t){t.updateTime=(new Date).getTime(),e.info=t,localStorage.setItem("account",Object(a["encrypt"])(JSON.stringify(t),"admio"))})),Object(s["a"])(o,"account.REMOVE",(function(e){e.info=null,localStorage.removeItem("account")})),o),l={};t["default"]={state:r,mutations:i,actions:l}},8:function(e,t){},"894f":function(e,t,n){"use strict";n.r(t);var o=n("2b0e"),s=n("c1df"),a=n.n(s);o["default"].prototype.$time=a.a,t["default"]=a.a},8988:function(e,t,n){},"8bb8":function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col",[n("el-form",{ref:"form",attrs:{model:e.form,"label-width":"80px"}},[n("el-form-item",{attrs:{prop:"title",label:"应用名称",rules:{required:!0,message:"不能为空",trigger:"blur"}}},[n("el-input",{attrs:{disabled:e.loading},model:{value:e.form.title,callback:function(t){e.$set(e.form,"title",t)},expression:"form.title"}})],1),n("el-form-item",{attrs:{label:"应用简介"}},[n("el-input",{attrs:{disabled:e.loading,type:"textarea"},model:{value:e.form.desc,callback:function(t){e.$set(e.form,"desc",t)},expression:"form.desc"}})],1),n("el-form-item",{attrs:{label:"是否公开"}},[n("el-radio-group",{attrs:{disabled:e.loading},model:{value:e.form.own,callback:function(t){e.$set(e.form,"own",t)},expression:"form.own"}},[n("el-radio",{attrs:{label:1,disabled:""}},[e._v("私有")]),n("el-radio",{attrs:{label:0}},[e._v("公开")])],1)],1),n("el-form-item",{attrs:{label:"容量大小",prop:"size",rules:[{required:!0,message:"不能为空",trigger:"blur"}]}},[n("el-input",{attrs:{disabled:e.loading},model:{value:e.form.size,callback:function(t){e.$set(e.form,"size",e._n(t))},expression:"form.size"}},[n("template",{slot:"append"},[e._v("GB")])],2)],1),n("el-form-item",[n("el-button",{attrs:{type:"primary",loading:e.loading},on:{click:function(t){return e.onSubmit("form")}}},[e._v("创建")])],1)],1)],1)],1)},s=[],a={name:"add",title:"添加应用",icon:"iconfont icon-tianjiayingyong",index:0,data:function(){return{form:{title:"",desc:"",own:0,size:1},loading:!1}},methods:{onSubmit:function(e){var t=this;this.loading=!0,this.$refs[e].validate((function(e){if(e){var n=JSON.parse(JSON.stringify(t.form));n.uuid=t.$store.state.user.info._id,n.secret=t.$uuid(),n.createTime=(new Date).getTime(),n.updateTime=(new Date).getTime(),t.$http("apps/insert",{data:n}).then((function(){t.$message.success("添加应用 ".concat(n.title," 成功!")),t.loading=!1,t.$router.push({path:"/apps"})})).catch((function(e){t.$message.error(e),t.loading=!1}))}else t.$message.error(e),t.loading=!1}))}}},r=a,i=n("2877"),l=Object(i["a"])(r,o,s,!1,null,"25534d4a",null);t["default"]=l.exports},9:function(e,t){},b5c4:function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col")],1)},s=[],a={name:"account",title:"我的信息",icon:"iconfont icon-yonghu",index:1,show:!0},r=a,i=n("2877"),l=Object(i["a"])(r,o,s,!1,null,"9461501e",null);t["default"]=l.exports},be3b:function(e,t,n){"use strict";n.r(t);var o=n("2b0e"),s=n("bc3a"),a=n.n(s),r=n("0009");a.a.defaults.headers.post["Authorization"]="94f3eee0-218f-41fc-9318-94cf5430fc7f";var i="/";o["default"].prototype.$http=function(e,t){return new Promise((function(n,o){Object(r["default"])("[DATA] (input) ->",e,JSON.stringify(t)),a.a.post("".concat(i).concat(e),t).then((function(e){Object(r["default"])("[DATA] (output) ->",JSON.stringify(e.data)),e.data.success?n(e.data.data):o(e.data.data)})).catch((function(e){Object(r["default"])("[DATA] (error) ->",e),o(e)}))}))},t["default"]=o["default"]},be72:function(e,t,n){"use strict";n.r(t);var o=n("2b0e"),s=n("6821f"),a=n.n(s);o["default"].prototype.$md5=a.a,t["default"]=o["default"]},be9b:function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col",[n("el-tree",{attrs:{props:e.database,load:e.loadNode,lazy:"","show-checkbox":""},on:{"check-change":e.handleCheckChange},scopedSlots:e._u([{key:"default",fn:function(t){var o=t.node,s=t.data;return n("span",{staticClass:"custom-tree-node"},[n("span",[e._v(e._s(o.label))]),n("span",[n("el-button",{attrs:{type:"text",size:"mini"},on:{click:function(){return e.append(s)}}},[e._v("Append")]),n("el-button",{attrs:{type:"text",size:"mini"},on:{click:function(){return e.remove(o,s)}}},[e._v("Delete")])],1)])}}])})],1)],1)},s=[],a={name:"info",title:"数据库管理",icon:"iconfont icon-database1",index:1,data:function(){return{database:[]}},methods:{loadNode:function(e,t){this.console(e),t([])},handleCheckChange:function(){}}},r=a,i=n("2877"),l=Object(i["a"])(r,o,s,!1,null,"6b1fe03b",null);t["default"]=l.exports},c321:function(e,t,n){var o={"./account.vue":"b5c4","./account/password.vue":"185b","./app/add.vue":"8bb8","./app/database.vue":"be9b","./app/edit/#id.vue":"da44","./app/info/#id.vue":"2948","./apps.vue":"45b4","./home.vue":"6511","./info.vue":"13ac","./login.vue":"dd7b","./reg.vue":"4649"};function s(e){var t=a(e);return n(t)}function a(e){if(!n.o(o,e)){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t}return o[e]}s.keys=function(){return Object.keys(o)},s.resolve=a,e.exports=s,s.id="c321"},c472:function(e,t,n){},c955:function(e,t,n){"use strict";n("288b")},d2d8:function(e,t,n){"use strict";n("02bb")},d815:function(e,t,n){},d84b:function(e,t,n){"use strict";n("d815")},da44:function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",[n("el-col",[n("el-form",{ref:"form",attrs:{model:e.form,disable:e.loading,"label-width":"80px"}},[n("el-form-item",{attrs:{prop:"title",label:"应用名称",rules:{required:!0,message:"不能为空",trigger:"blur"}}},[n("el-input",{attrs:{disabled:e.loading},model:{value:e.form.title,callback:function(t){e.$set(e.form,"title",t)},expression:"form.title"}})],1),n("el-form-item",{attrs:{label:"应用简介"}},[n("el-input",{attrs:{disabled:e.loading,type:"textarea"},model:{value:e.form.desc,callback:function(t){e.$set(e.form,"desc",t)},expression:"form.desc"}})],1),n("el-form-item",{attrs:{label:"是否公开"}},[n("el-radio-group",{attrs:{disabled:e.loading},model:{value:e.form.own,callback:function(t){e.$set(e.form,"own",t)},expression:"form.own"}},[n("el-radio",{attrs:{label:1,disabled:""}},[e._v("私有")]),n("el-radio",{attrs:{label:0}},[e._v("公开")])],1)],1),n("el-form-item",{attrs:{label:"容量大小",prop:"size",rules:[{required:!0,message:"不能为空",trigger:"blur"}]}},[n("el-input",{attrs:{disabled:e.loading},model:{value:e.form.size,callback:function(t){e.$set(e.form,"size",e._n(t))},expression:"form.size"}},[n("template",{slot:"append"},[e._v("GB")])],2)],1),n("el-form-item",[n("el-button",{attrs:{type:"primary",loading:e.loading},on:{click:function(t){return e.onSubmit("form")}}},[e._v("保存")])],1)],1)],1)],1)},s=[],a={name:"edit",title:"修改应用",icon:"iconfont icon-tianjiayingyong",index:0,data:function(){return{form:{title:"",desc:"",own:0,size:1},loading:!1}},methods:{onSubmit:function(e){var t=this;this.loading=!0,this.$refs[e].validate((function(e){if(e){var n=JSON.parse(JSON.stringify(t.form));n.updateTime=(new Date).getTime(),t.console(n),t.$http("apps/update",{data:n,where:{_id:t.$route.params.id,uuid:t.$store.state.user.info._id}}).then((function(){t.$message.success("保存成功!"),t.loading=!1})).catch((function(e){t.$message.error(e),t.loading=!1}))}else t.$message.error(e),t.loading=!1}))}},created:function(){var e=this;this.loading=!0,this.$http("apps/once",{where:{_id:this.$route.params.id,uuid:this.$store.state.user.info._id},other:{show:{_id:0,uuid:0,createTime:0,secret:0}}}).then((function(t){t?e.form=t:(e.$message.error("没有此应用存在!"),e.$router.push({path:"/app"})),e.loading=!1})).catch((function(t){e.$message.error(t)}))}},r=a,i=n("2877"),l=Object(i["a"])(r,o,s,!1,null,"62828399",null);t["default"]=l.exports},dd7b:function(e,t,n){"use strict";n.r(t);var o=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-row",{staticClass:"main",attrs:{type:"flex",justify:"center"}},[n("el-col",{staticStyle:{"min-width":"300px","max-width":"400px"},attrs:{span:12}},[n("el-row",{staticClass:"form"},[n("el-col",[n("label",{class:{focus:e.focus.username}},[n("span",[e._v("用户名")]),n("el-input",{ref:"username",attrs:{autofocus:!0,placeholder:e.focus.username?e.placeholder.username:"",disabled:e.loading},on:{focus:function(t){e.focus.username=!0},blur:function(t){e.focus.username=""!==e.form.username}},model:{value:e.form.username,callback:function(t){e.$set(e.form,"username",t)},expression:"form.username"}})],1)]),n("el-col",[n("label",{class:{focus:e.focus.password}},[n("span",[e._v("密码")]),n("el-input",{ref:"password",attrs:{"show-password":"",placeholder:e.focus.password?e.placeholder.password:"",disabled:e.loading},on:{focus:function(t){e.focus.password=!0},blur:function(t){e.focus.password=""!==e.form.password}},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.login.apply(null,arguments)}},model:{value:e.form.password,callback:function(t){e.$set(e.form,"password",t)},expression:"form.password"}})],1)]),n("el-col",[n("el-row",{attrs:{type:"flex",align:"bottom"}},[n("el-col",{attrs:{span:12}},[n("el-button",{attrs:{type:"primary",loading:e.loading},on:{click:e.login}},[e._v("登陆")])],1),n("el-col",{staticClass:"text-right",attrs:{span:12}},[n("router-link",{attrs:{to:"/reg"}},[e._v("注册一个账号")])],1)],1)],1)],1)],1)],1)},s=[],a={name:"login",title:"登陆帐号",data:function(){return{focus:{username:!1,password:!1},placeholder:{username:"手机号/邮箱",password:"格式:[A-Z]{1}(.*){7,}"},form:{username:"",password:""},loading:!1}},methods:{login:function(){var e=this;if(""===this.form.username)return this.$refs.username.focus();if(""===this.form.password)return this.$refs.password.focus();this.loading=!0;var t=JSON.parse(JSON.stringify(this.form));t.password=this.$md5(t.password),this.$http("users/once",{where:t}).then((function(n){n?(e.$message.success("".concat(t.username," 用户登陆成功!")),e.$store.commit("account.UPDATE",n),e.$router.push({path:e.$route.query.redirect||"/home"})):e.$message.error("".concat(t.username," 用户登陆失败,请稍后再试!")),e.loading=!1})).catch((function(t){e.$message.error("未知错误: ".concat(JSON.stringify(t))),e.loading=!1}))}},mounted:function(){}},r=a,i=(n("c955"),n("2877")),l=Object(i["a"])(r,o,s,!1,null,"5f562552",null);t["default"]=l.exports},e0ac:function(e,t,n){"use strict";var o=n("1c46");function s(e,t){var n=o.createCipher("aes192",t),s=n.update(e,"utf8","hex");return s+=n.final("hex"),s}function a(e,t){var n=o.createDecipher("aes192",t),s=n.update(e,"hex","utf8");return s+=n.final("utf8"),s}function r(e){var t=o.createHash("md5");return t.update(e),e=t.digest("hex"),e}e.exports={encrypt:s,decrypt:a,md5:r}},e75c:function(e,t,n){var o={"./user.js":"7f34"};function s(e){var t=a(e);return n(t)}function a(e){if(!n.o(o,e)){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t}return o[e]}s.keys=function(){return Object.keys(o)},s.resolve=a,e.exports=s,s.id="e75c"},f666:function(e,t,n){var o={"./ElementUI.js":"1b03","./axios.js":"be3b","./console.js":"0009","./debug.js":"1cee","./md5.js":"be72","./moment.js":"894f","./uuid.js":"1179"};function s(e){var t=a(e);return n(t)}function a(e){if(!n.o(o,e)){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t}return o[e]}s.keys=function(){return Object.keys(o)},s.resolve=a,e.exports=s,s.id="f666"}});
//# sourceMappingURL=app.bedf3c76.js.map