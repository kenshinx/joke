$(document).ready(function(){

	$("#create-record").click(function(){

		entry = "\
		<tr class='entry-edit'>\
			<td><input type='text' name='domain'></input></td>\
			<td><input type='text' name='ip'></input></td>\
			<td class='action'>\
				<button id='save-record' type='submit' class='btn btn-success btn-small' onclick='save_edit(this)'>\
				<i class='icon-plus'></i> Save</button>\
				<button id='remove-edit' class='btn btn-link btn-small' onclick='remove_edit(this)' >\
				<i class='icon-remove'></i> Remove</button>\
			</td>\
		</tr>\
		"
		$("#record-add").prepend(entry);
	});


})


function remove_edit(btn){
	$(btn).parent().parent().remove();
}

function save_edit(btn){
	tr = $(btn).parent().parent();
	domain = tr.find("input[name='domain']").val();
	ip = tr.find("input[name='ip']").val();
	$.ajax({
		url: "",
		type: "POST",
		data: {domain:domain,ip:ip},
		success: function(){
					entry = '<tr class="entry"><td id="domain">'
					entry += domain
					entry +='</td><td id="ip">'
					entry += ip
					entry += '</td><td class="action">\
			                        <button id="edit-record" class="btn btn-info btn-small" onclick="edit_record(this)">\
			                            <i class="icon-edit"></i> Edit\
			                        </button>\
			                        <button id="del-record" class="btn btn-danger btn-small" onclick="del_record(this)">\
			                            <i class="icon-remove"></i> Delete\
			                        </button>\
			                    </td>\
			                </tr>'
					$("#record-list").prepend(entry);
					tr.remove()
				},
		error: function(xhr,status,err){
			alert(xhr.responseText);
		}
	})

}


function del_record(btn){
	bootbox.confirm("Confirm Detete?", function(confirm) {
		if (!confirm) {
			return;
		}

		tr = $(btn).parent().parent();
		domain = tr.find("td#domain").html();
		ip = tr.find("td#ip").html();
		$.ajax({
			type : "POST",
			url : "/dns/del",
			data: {"domain":domain,ip:ip},
			success : function(){
				tr.remove();
			},
			error: function(xhr,status,err){
				alert(xhr.responseText);
			}

		})
	})
}

function edit_record(btn){
	tr = $(btn).parent().parent();
	domain = tr.find("td#domain").html();
	ip = tr.find("td#ip").html();

	raw = tr.html()

	entry = "<td><input type='text' name='domain' value='"
	entry += domain
	entry += "' disabled></input></td> <td><input type='text' name='ip' value='"
	entry += ip
	entry += "'></input></td>\
				<td class='action'>\
					<button id='save-record' type='submit' class='btn btn-success btn-small' onclick='save_edit(this)'>\
					<i class='icon-plus'></i> Save</button>\
					<button id='cancel-edit' class='btn btn-warning btn-small' onclick='cancel_edit(this,raw)' >\
					<i class='icon-minus'></i> Cancel</button>\
				</td>\
	"
	tr.html(entry);
}


function cancel_edit(btn,raw_html){
	tr = $(btn).parent().parent();
	tr.html(raw_html);
}

