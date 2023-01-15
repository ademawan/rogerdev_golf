
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
      
        ajax: "/dashboard/master/daftar-teknisi",
        columns: [
          {data: 'DT_RowIndex', name: 'DT_RowIndex'},
          { data: 'nama_lengkap', name: 'nama_lengkap', class:'text-center' },
          { data: 'jabatan', name: 'jabatan' },
          { data: 'status', name: 'status' },
          { data: 'action', name: 'action' },
            ]
      });
      

      $('#tombolTambahTeknisi').click(function(){
        $('#saveBtnTambahTeknisi').addClass('show');
        $('#saveBtnUpdateTeknisi').attr('hidden', 'true'); 
        $('#saveBtnTambahTeknisi').removeAttr('hidden'); 
  
      });

  
      // tambah teknisi
      $('#saveBtnTambahTeknisi').click( function(e){
  
        e.preventDefault();
        $(this).html('Sending..');
    
        $.ajax({
          data: $('#formtambahteknisi').serialize(),
          url: "/dashboard/master/daftar-teknisi",
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
                      $('#saveBtnTambahTeknisi').html('Tambah');
                    });
                  }else{
                    $('#formtambahteknisi').trigger("reset");
                    table.draw();
                    $('#saveBtnTambahTeknisi').html('Tambah');
                    $('#tombolTambahTeknisi').trigger('click');
                    $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses tambah data !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                    $('.alert').delay(2000).fadeOut(function() {
                      $(this).remove(); 
                   });
                  }
  
         
          },
          
       });
  
      })

      //edit
    $('body').on('click', '.editTeknisi', function () {

      $('#saveBtnTambahTeknisi').attr('hidden', 'true'); 
      $('#saveBtnUpdateTeknisi').removeAttr('hidden'); 
      $('#saveBtnUpdateTeknisi').html('Update');

      var teknisi_id = $(this).data('id');
      $.get("/dashboard/master/daftar-teknisi" +'/' + teknisi_id +'/edit', function (data) {
        $('#tambahteknisi').addClass('show');
          
          $('#id_teknisi').val(data.id);
          $('#nama_lengkap').val(data.nama_lengkap);
          $('#jabatan').val(data.jabatan);
          $('#status').val(data.status);
          $('#qty').val(data.qty);
          
          $("html, body").animate({ scrollTop: 0 }, "slow");    
          // $('#scroll-ku').scrollTop = 0;

          


          
      })
   });


   //update
   $('#saveBtnUpdateTeknisi').click(function (e) {

    e.preventDefault();
    $(this).html('Sending..');
    var id_teknisi = $('#id_teknisi').val();
    $.ajax({
      data: $('#formtambahteknisi').serialize(),
      url: "/dashboard/master/daftar-teknisi"+'/'+id_teknisi,
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
                  $('#saveBtnUpdateTeknisi').html('Update');
                });
              }else{
                $(this).html('Update');
                $('#formtambahteknisi').trigger("reset");
                table.draw();
                $('#saveBtnTambahTeknisi').html('Tambah');
                $('#tombolTambahTeknisi').trigger('click');   
                $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses Update !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                $('.alert').delay(2000).fadeOut(function() {
                  $(this).remove(); 
               });
              }

     
      },
      
    });

  })


  // Batal
   $('#saveBtnBatalTeknisi').click( function(e){
     $(this).html('Sending..');
     e.preventDefault();
        $('#formtambahteknisi').trigger("reset");
        $('#saveBtnBatalTeknisi').html('Batal');
        $('#tombolTambahTeknisi').trigger('click');
        $(document).find('span.error-text').text('');


   })


    //delete
    $('body').on('click', '.deleteTeknisi', function () {
     
        var teknisi_id = $(this).data("id");
        var r = confirm("Are You sure want to delete !");
         if (r == true) {
           $.ajax({
            type: "DELETE",
            url: "/dashboard/master/daftar-teknisi"+'/'+teknisi_id,
            success: function (data) {
              $('.table-responsive').before(`<div class="alert alert-danger alert-dismissible fade show" role="alert"> Sukses hapus teknisi!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
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