$(document).ready(function(){


	$("#create-record").click(function(){

		entry = "\
		<tr class='entry-edit'>\
			<td><input type='text' name='domain'></input></td>\
			<td><input type='text' name='ip'></input></td>\
			<td class='action'>\
				<button id='save-record' type='submit' class='btn btn-success btn-small' onclick='save_edit(this)'>\
				<i class='icon-plus'></i> Save</button>\
				<button id='cancel-edit' class='btn btn-warning btn-small' onclick='cancel_edit(this)' >\
				<i class='icon-minus'></i> Cancel</button>\
			</td>\
		</tr>\
		"
		$("#record-add").prepend(entry);
	});


})


function cancel_edit(entry){
	$(entry).parent().parent().remove();
}

function save_edit(entry){
	p = $(entry).parent().parent();
	domain = p.find("input[name='domain']").val();
	ip = p.find("input[name='ip']").val();
	$.post("", {domain:domain,ip:ip})
	.done(function(){
		entry = '<tr class="entry"><td>'
		entry += domain
		entry +='</td><td>'
		entry += ip
		entry += '</td><td class="action">\
                        <button id="edit-record" class="btn btn-info btn-small">\
                            <i class="icon-edit"></i> Edit\
                        </button>\
                        <button id="del-record" class="btn btn-danger btn-small">\
                            <i class="icon-remove"></i> Delete\
                        </button>\
                    </td>\
                </tr>'
		$("#record-list").prepend(entry);
		p.remove()
	})
	.fail(function(){
		alert("failed");
	})

}

