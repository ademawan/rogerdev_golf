$(function () {
  
  $.ajaxSetup({
      headers: {
          'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
      }
    });

   var table= $('#myTable').DataTable( {
      
      processing: true,
      serverSide: true,
      paging:true,
      scrollX: true,
      scrollY: true,
    
      ajax: "/dashboard/master/daftar-user",
      columns: [
              { data: 'nama_lengkap', name: 'nama_lengkap', class:'text-center' },
              { data: 'username', name: 'username' },
              { data: 'email', name: 'email', orderable:false },
              { data: 'jabatan', name: 'jabatan', orderable:false },
              { data: 'status', name: 'status', orderable:false },
              { data: 'role', name: 'role', orderable:false },
              { data: 'action', name: 'action', orderable:false },
          ]
    });
    

    $('#tombolRefreshTable').click(function(){
      table.draw();
    });
    
    $('#tombolTambahUser').click(function(){
      $('#saveBtnTambahUser').addClass('show');
      $('#saveBtnUpdateUser').attr('hidden', 'true'); 
      $('#saveBtnTambahUser').removeAttr('hidden'); 
    });


    // Tambah
    $('#saveBtnTambahUser').click( function(e){

      e.preventDefault();
      $(this).html('Sending..');
  
      $.ajax({
        data: $('#formtambahuser').serialize(),
        url: "/dashboard/master/daftar-user",
        type: "POST",
        dataType: 'json',
        beforeSend:function(){
              $(document).find('span.error-text').text('');
            },
        success: function (data) {

          
                if(data.status==0){
                  $.each(data.error, function(frefix, val){
                    $('span.'+frefix+'_error').text(val[0]);
                    $('span.'+frefix+'_errornotif').text('['+val[0]+']');
                    $('#saveBtnTambahUser').html('Tambah');
                  });
                }else{
                  $('#formtambahuser').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahUser').html('Tambah');
                  $('#tombolTambahUser').trigger('click');
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses tambah data !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                    $(this).remove(); 
                 });
                    
                }
        },
        
      });

    })

   

    //edit
    $('body').on('click', '.editUser', function () {
      $('#saveBtnTambahUser').attr('hidden', 'true'); 
      $('#saveBtnUpdateUser').removeAttr('hidden'); 
      $('#saveBtnUpdateUser').html('Update');




      var user_id = $(this).data('id');
      $.get("/dashboard/master/daftar-user" +'/' + user_id +'/edit', function (data) {
        $('#tambahuser').addClass('show');
        $('#colpassword').attr('hidden', 'true');

          $('#id_user').val(data.id);
          $('#nama_lengkap').val(data.nama_lengkap);
          $('#username').val(data.username);
          $('#email').val(data.email);
          $('#jabatan').val(data.jabatan);
          $('#status').val(data.status);
          $('#role').val(data.role_id);

          $("html, body").animate({ scrollTop: 0 }, "slow");
          // $('#scroll-ku').scrollTop = 0;
          
      })
   });


    //update
    $('#saveBtnUpdateUser').click(function (e) {

      e.preventDefault();
      $(this).html('Sending..');
      var id_user = $('#id_user').val();
      $.ajax({
        data: $('#formtambahuser').serialize(),
        url: "/dashboard/master/daftar-user"+'/'+id_user,
        type: "PUT",
        dataType: 'json',
        beforeSend:function(){
              $(document).find('span.error-text').text('');
            },
        success: function (data) {

          
                if(data.status==0){
                  $.each(data.error, function(frefix, val){
                    $('span.'+frefix+'_error').text(val[0]);
                    $('span.'+frefix+'_errornotif').text('['+val[0]+']');
                    $('#updateBtnUpdateUser').html('Update');
                  });
                }else{
                  $(this).html('Update');
                  $('#formtambahuser').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahUser').html('Tambah');
                  $('#colpassword').removeAttr('hidden');

                  $('#tombolTambahUser').trigger('click');   
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses Update !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                     $(this).remove(); 
                  });
                }
        },
        
      });

    })


    //batal
   $('#saveBtnBatalUser').click( function(e){
     $(this).html('Sending..');
     e.preventDefault();
     $('#formtambahuser').trigger("reset");
     $('#colpassword').removeAttr('hidden');

        $('#saveBtnBatalUser').html('Batal');
        $('#tombolTambahUser').trigger('click');
        $(document).find('span.error-text').text('');

   })


    //delete
    $('body').on('click', '.deleteUser', function () {
     
        var user_id = $(this).data("id");
        var r = confirm("Are You sure want to delete !");
         if (r == true) {
           $.ajax({
            type: "DELETE",
            url: "/dashboard/master/daftar-user"+'/'+user_id,
            success: function (data) {
              $('.table-responsive').before(`<div class="alert alert-danger alert-dismissible fade show" role="alert"> Sukses hapus user!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
              table.draw();
              $('.alert').delay(2000).fadeOut(function() {
                $(this).remove(); 
             });
            },
            error: function (data) {
                console.log('Error:', data);
            }
           });
         } else {
           table.draw();
         }
    });

   

 } );