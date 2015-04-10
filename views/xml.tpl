<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
 
 
<style type="text/css">
         *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

	
	   #fm{
            margin:0;
            padding:10px 30px;
        }
        .ftitle{
            font-size:14px;
            font-weight:bold;
            padding:5px 0;
            margin-bottom:10px;
            border-bottom:1px solid #ccc;
        }
        .fitem{
            margin-bottom:5px;
        }
        .fitem label{
            display:inline-block;
            width:80px;
        }
        .fitem input{
            width:160px;
        }
		
		.fieldspan{ width:100%; clear:both ; display:block ;height:25px;  }
		.fieldspan .formlabel{width:90px; float:left ; }
		.fieldspan .forminput{width:390px; float:left ; }
    </style>
	<link rel="stylesheet" type="text/css" href="/static/jui/themes/default/easyui.css">
	<link rel="stylesheet" type="text/css" href="/static/jui/themes/icon.css">

	<script type="text/javascript" src="/static/jui/jquery.min.js"></script>
	<script type="text/javascript" src="/static/jui/jquery.easyui.min.js"></script>
	<script type="text/javascript" src="/static/jui/locale/easyui-lang-zh_CN.js"></script>




    <script type="text/javascript">
        var url;
        function newUser(){
            $('#dlg').dialog('open').dialog('setTitle','新建{{.Xmod.Label}}');
            $('#fm').form('clear');
            url = '/xmldata_save/?FormModName={{.Xmod.Name}}';
        }
        function editUser(){
            var row = $('#dg').datagrid('getSelected');
            if (row){
                $('#dlg').dialog('open').dialog('setTitle','编辑{{.Xmod.Label}}');
                $('#fm').form('load',row);
                url = '/xmldata_save/?FormModName={{.Xmod.Name}}';
            }
        }
        function saveUser(){
            $('#fm').form('submit',{
                url: url,
                onSubmit: function(){
                    return $(this).form('validate');
                },
                success: function(result){
                    var result = eval('('+result+')');
                    if (result.errorMsg){
                        $.messager.show({
                            title: 'Error',
                            msg: result.errorMsg
                        });
                    } else {
                        $('#dlg').dialog('close');        // close the dialog
                        $('#dg').datagrid('reload');    // reload the user data
                    }
                }
            });
        }
        function destroyUser(){
            var row = $('#dg').datagrid('getSelected');
            if (row){
                $.messager.confirm('Confirm','Are you sure you want to destroy this user?',function(r){
                    if (r){
                        $.post('/xmldata_save/?FormModName={{.Xmod.Name}}',{Deleteid:row.id},function(result){
                            if (result.success){
                                $('#dg').datagrid('reload');    // reload the user data
                            } else {
                                $.messager.show({    // show error message
                                    title: 'Error',
                                    msg: result.errorMsg
                                });
                            }
                        },'json');
                    }
                });
            }
        }
		
		
    </script>

</head>

<body>


		
    <table id="dg" title="{{.Xmod.Label}}" class="easyui-datagrid" style="width:800px;height:450px"
            url="/xmldata_get/?FormModName={{.Xmod.Name}}" toolbar="#toolbar" pagination="true"   rownumbers="true" fitColumns="true" singleSelect="true">
        <thead>
	         <tr>
			 
				
				 {{.Xmod | renderformxmllistdq}}
            </tr>
        </thead>
    </table>
	
    <div id="toolbar">
        <a href="javascript:void(0)" class="easyui-linkbutton" iconCls="icon-add" plain="true" onclick="newUser()">新建</a>
        <a href="javascript:void(0)" class="easyui-linkbutton" iconCls="icon-edit" plain="true" onclick="editUser()">编辑</a>
        <a href="javascript:void(0)" class="easyui-linkbutton" iconCls="icon-remove" plain="true" onclick="destroyUser()">删除</a>
    </div>
    
    <div id="dlg" class="easyui-dialog" style="width:700px;height:580px;padding:10px 20px"  closed="true" buttons="#dlg-buttons">
        <div class="ftitle">{{.Xmod.Label}}信息</div>
        <form id="fm" method="post" novalidate>
		<input type="hidden" value="" id="id" name="id" />
           {{.Xmod | renderformxmldq}}
 		</form>
    </div>
    <div id="dlg-buttons">
        <a href="javascript:void(0)" class="easyui-linkbutton c6" iconCls="icon-ok" onclick="saveUser()" style="width:90px">保存</a>
        <a href="javascript:void(0)" class="easyui-linkbutton" iconCls="icon-cancel" onclick="javascript:$('#dlg').dialog('close')" style="width:90px">取消</a>
    </div>
	
<script>
$(".formlabelhide").parent().hide();
</script>
</body>
</html>
