(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-4e6299a1"],{a434:function(e,t,a){"use strict";var l=a("23e7"),o=a("23cb"),n=a("5926"),i=a("07fa"),s=a("7b0b"),r=a("65f0"),d=a("8418"),u=a("1dde"),m=u("splice"),c=Math.max,f=Math.min,p=9007199254740991,b="Maximum allowed length exceeded";l({target:"Array",proto:!0,forced:!m},{splice:function(e,t){var a,l,u,m,h,v,g=s(this),y=i(g),x=o(e,y),k=arguments.length;if(0===k?a=l=0:1===k?(a=0,l=y-x):(a=k-2,l=f(c(n(t),0),y-x)),y+a-l>p)throw TypeError(b);for(u=r(g,l),m=0;m<l;m++)h=x+m,h in g&&d(u,m,g[h]);if(u.length=l,a<l){for(m=x;m<y-l;m++)h=m+l,v=m+a,h in g?g[v]=g[h]:delete g[v];for(m=y;m>y-l+a;m--)delete g[m-1]}else if(a>l)for(m=y-l;m>x;m--)h=m+l-1,v=m+a-1,h in g?g[v]=g[h]:delete g[v];for(m=0;m<a;m++)g[m+x]=arguments[m+2];return g.length=y-l+a,u}})},bcf9:function(e,t,a){"use strict";a("f0f3")},ea99:function(e,t,a){"use strict";a.r(t);var l=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("el-card",{staticClass:"box-card"},[a("el-button",{attrs:{type:"primary"},on:{click:function(t){e.dialogVisible=!0}}},[e._v("添加学生")]),a("el-table",{staticStyle:{width:"100%"},attrs:{data:e.tableData}},[a("el-table-column",{attrs:{type:"index",width:"50",label:"序号"}}),a("el-table-column",{attrs:{prop:"name",label:"学生姓名"}}),a("el-table-column",{attrs:{prop:"createAt",label:"创建时间"}}),a("el-table-column",{attrs:{prop:"sex",label:"性别"}}),a("el-table-column",{attrs:{prop:"professional",label:"班级"}}),a("el-table-column",{attrs:{prop:"phone",label:"手机号"}}),a("el-table-column",{attrs:{prop:"group",label:"分组"}}),a("el-table-column",{attrs:{prop:"seat",label:"座位"}}),a("el-table-column",{attrs:{label:"操作",width:"150"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-button",{attrs:{size:"mini"},on:{click:function(a){return e.handleEdit(t.$index,t.row)}}},[e._v("编辑")]),a("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(a){return e.handleDelete(t.$index,t.row)}}},[e._v("删除")])]}}])})],1)],1),a("el-dialog",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],attrs:{title:"添加学生",visible:e.dialogVisible,width:"30%"},on:{"update:visible":function(t){e.dialogVisible=t}}},[a("el-form",{attrs:{"label-width":"80px"}},[a("el-form-item",{attrs:{label:"学生名称"}},[a("el-input",{model:{value:e.studentName,callback:function(t){e.studentName=t},expression:"studentName"}})],1),a("el-form-item",{attrs:{label:"学生性别"}},[a("el-radio-group",{model:{value:e.studentSex,callback:function(t){e.studentSex=t},expression:"studentSex"}},[a("el-radio",{attrs:{label:"男"}}),a("el-radio",{attrs:{label:"女"}})],1)],1),a("el-form-item",{attrs:{label:"学生专业"}},[a("el-select",{attrs:{placeholder:"请选择活动区域"},model:{value:e.studentProfessional,callback:function(t){e.studentProfessional=t},expression:"studentProfessional"}},[a("el-option",{attrs:{label:"软件技术",value:"软件技术"}}),a("el-option",{attrs:{label:"人力资源",value:"人力资源"}}),a("el-option",{attrs:{label:"会计",value:"会计"}}),a("el-option",{attrs:{label:"艺术",value:"艺术"}}),a("el-option",{attrs:{label:"机器人",value:"机器人"}}),a("el-option",{attrs:{label:"体育",value:"体育"}})],1)],1),a("el-form-item",{attrs:{label:"手机号码"}},[a("el-input",{model:{value:e.studentPhone,callback:function(t){e.studentPhone=t},expression:"studentPhone"}})],1),a("el-form-item",{attrs:{label:"所属分组"}},[a("el-select",{attrs:{placeholder:"请选择活动区域"},model:{value:e.studentGroup,callback:function(t){e.studentGroup=t},expression:"studentGroup"}},[a("el-option",{attrs:{label:"一",value:"G1"}}),a("el-option",{attrs:{label:"二",value:"G2"}}),a("el-option",{attrs:{label:"三",value:"G3"}}),a("el-option",{attrs:{label:"四",value:"G4"}})],1)],1),a("el-form-item",{attrs:{label:"座位号"}},[a("el-input",{model:{value:e.studentSeat,callback:function(t){e.studentSeat=t},expression:"studentSeat"}})],1)],1),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(t){e.dialogVisible=!1}}},[e._v("取 消")]),a("el-button",{attrs:{type:"primary"},on:{click:e.commitStudentHandler}},[e._v("确 定")])],1)],1),a("el-dialog",{directives:[{name:"loading",rawName:"v-loading",value:e.loading,expression:"loading"}],attrs:{title:"编辑学生",visible:e.modifyDialog,width:"30%"},on:{"update:visible":function(t){e.modifyDialog=t}}},[a("el-form",{attrs:{"label-width":"150px"}},[a("el-form-item",{attrs:{label:"学生名称"}},[a("el-input",{model:{value:e.modifyForm.name,callback:function(t){e.$set(e.modifyForm,"name",t)},expression:"modifyForm.name"}})],1),a("el-form-item",{attrs:{label:"学生性别"}},[a("el-radio-group",{model:{value:e.modifyForm.sex,callback:function(t){e.$set(e.modifyForm,"sex",t)},expression:"modifyForm.sex"}},[a("el-radio",{attrs:{label:"男"}}),a("el-radio",{attrs:{label:"女"}})],1)],1),a("el-form-item",{attrs:{label:"学生专业"}},[a("el-select",{attrs:{placeholder:"请选择活动区域"},model:{value:e.modifyForm.professional,callback:function(t){e.$set(e.modifyForm,"professional",t)},expression:"modifyForm.professional"}},[a("el-option",{attrs:{label:"软件技术",value:"软件技术"}}),a("el-option",{attrs:{label:"人力资源",value:"人力资源"}}),a("el-option",{attrs:{label:"会计",value:"会计"}}),a("el-option",{attrs:{label:"艺术",value:"艺术"}}),a("el-option",{attrs:{label:"机器人",value:"机器人"}})],1)],1),a("el-form-item",{attrs:{label:"手机号码"}},[a("el-input",{model:{value:e.modifyForm.phone,callback:function(t){e.$set(e.modifyForm,"phone",t)},expression:"modifyForm.phone"}})],1),a("el-form-item",{attrs:{label:"所属分组"}},[a("el-select",{attrs:{placeholder:"请选择活动区域"},model:{value:e.modifyForm.group,callback:function(t){e.$set(e.modifyForm,"group",t)},expression:"modifyForm.group"}},[a("el-option",{attrs:{label:"一",value:"G1"}}),a("el-option",{attrs:{label:"二",value:"G2"}}),a("el-option",{attrs:{label:"三",value:"G3"}}),a("el-option",{attrs:{label:"四",value:"G4"}})],1)],1),a("el-form-item",{attrs:{label:"座位号"}},[a("el-input",{model:{value:e.modifyForm.seat,callback:function(t){e.$set(e.modifyForm,"seat",t)},expression:"modifyForm.seat"}})],1)],1),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(t){e.modifyDialog=!1}}},[e._v("取 消")]),a("el-button",{attrs:{type:"primary"},on:{click:e.modifyStudentHandler}},[e._v("确 定")])],1)],1)],1)},o=[],n=(a("b0c0"),a("a434"),{name:"",data:function(){return{loading:!1,studentName:"",studentSex:"",studentProfessional:"",studentPhone:"",studentGroup:"",studentSeat:"",dialogVisible:!1,tableData:[],modifyForm:{},modifyDialog:!1,modifyIndex:0}},created:function(){this.loadData()},methods:{commitStudentHandler:function(){this.loading=!0;var e=this;this.$POST("/addstudent",{name:this.studentName,sex:this.studentSex,professional:this.studentProfessional,phone:this.studentPhone,group:this.studentGroup,seat:this.studentSeat},(function(t){e.loading=!1,200!=t.code?e.$message.error(t.message):(e.$message.success(t.message),e.dialogVisible=!1,e.tableData.push({id:t.data.id,name:t.data.name,createAt:t.data.createAt,sex:t.data.sex,professional:t.data.professional,phone:t.data.phone,group:t.data.group,seat:t.data.seat}))}))},loadData:function(){var e=this;this.$POST("/student",{},(function(t){console.log("student的vc：",t),200==t.code&&(e.tableData=t.data)}))},handleDelete:function(e,t){var a=this;console.log(e,t);var l=this;this.$confirm("你确定要删除该数据?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){a.$POST("/delstudent",{id:t.id},(function(t){200==t.code?l.tableData.splice(e,1):l.$message.error(t.message)}))})).catch((function(){a.$message({type:"info",message:"已取消删除"})}))},handleEdit:function(e,t){this.modifyIndex=e,this.modifyForm=Object.assign({},t),this.modifyDialog=!0},modifyStudentHandler:function(){var e=this;this.$POST("/modifystudent",this.modifyForm,(function(t){if(200==t.code){var a=JSON.parse(JSON.stringify(e.tableData));a[e.modifyIndex]=e.modifyForm,e.tableData=a,console.log(e.tableData),e.modifyDialog=!1}else e.$message.error(t.message)}))}}}),i=n,s=(a("bcf9"),a("2877")),r=Object(s["a"])(i,l,o,!1,null,"b47e07aa",null);t["default"]=r.exports},f0f3:function(e,t,a){}}]);
//# sourceMappingURL=chunk-4e6299a1.a8e5b526.js.map