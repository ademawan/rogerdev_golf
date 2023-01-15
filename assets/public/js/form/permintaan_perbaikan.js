$(function () {

  var container = document.querySelector('#container');
  var  input_sparepart;
  var input_qty;

  container.addEventListener('click', function(e){
    if(e.target.className == 'buttonCari'){
      input_sparepart = e.target.parentElement.previousElementSibling.previousElementSibling.firstElementChild;
      input_qty = e.target.parentElement.previousElementSibling.firstElementChild;
    }
    if(e.target.className == 'buttonHapusSparepart'){
      e.target.parentElement.parentElement.remove();
    }
  })


$('#buttonModalTambah').click(function(e){
  var dataId = $(this).parent().parent().find(':selected').val();
  var dataQty = $(this).parent().parent().find('#qty').val();
  input_sparepart.setAttribute('value', dataId);
  input_qty.setAttribute('value', dataQty);
  $('#buttonModalBatal').trigger('click');

});

//buttonTambahSparepart

$('.buttonTambahSparepart').click(function(e){
  
  $(this).parent().parent().parent().parent().append(` <div class="col-md-12">
  <div class="mb-3 row">
    <label for="inputtext" class="col-sm-3 col-form-label">Sparepart ID</label>
    <div class="col-sm-4 px-0" >
      <input type="text" name="sparepart[]" class="form-control inputsparepart" readonly>
        <span class="text-danger error-text sparepart_error"></span>
    </div>
    <div class="col-sm-3 px-1">
      <input type="text" name="qty[]" class="form-control" placeholder="QTY" style="display:block !important; padding:5px !important;" readonly>
        <span class="text-danger error-text qty_error"></span>
    </div>
    <div class="col-sm-1 px-0">
      <button type="button" class="buttonCari" style="display:block !important; padding:5px !important;" data-bs-toggle="modal" data-bs-target="#modalSparepart">
        Cari
      </button>
    </div>
    <div class="col-sm-1 px-0">
      <button type="button" class="buttonHapusSparepart" style="display:block !important; padding:5px !important;"> - </button>
    </div>
  </div>
</div>`);

});



//select2
  $('#selectjenismesin').select2();
  $('#selectteknisi').select2();

  $('#selectsparepart').select2({
    dropdownParent: $("#modalSparepart")
  });
  $('#tambahSparepartSS').click(function(e){
    e.preventDefault();
    $('#selectsparepart2').select2();
    $('#sparepart2').removeClass('visually-hidden');




  });


  
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
    
      ajax: "/dashboard/form/permintaan-perbaikan",
      columns: [
              { data: 'jenis_mesin', name: 'jenis_mesin', class:'text-center' },
              { data: 'teknisi', name: 'teknisi' },
              { data: 'user', name: 'user', orderable:false },
              { data: 'tipe_perbaikan', name: 'tipe_perbaikan', orderable:false },
              { data: 'masalah_kerusakan', name: 'masalah_kerusakan', orderable:false },
              { data: 'kronologis_kejadian', name: 'kronologis_kejadian', orderable:false },
              { data: 'analisa', name: 'analisa', orderable:false },
              { data: 'tindakan', name: 'tindakan', orderable:false },
              { data: 'progress', name: 'progress', orderable:false },
              { data: 'keterangan', name: 'keterangan', orderable:false },
              { data: 'mulai_terjadi', name: 'mulai_terjadi', orderable:false },
              { data: 'mulai_perbaikan', name: 'mulai_perbaikan', orderable:false },
              { data: 'selesai_perbaikan', name: 'selesai_perbaikan', orderable:false },
              { data: 'pemberi_order', name: 'pemberi_order', orderable:false },
          ],
          
    });
    

    $('#tombolRefreshTable').click(function(){
      table.draw();
    });
    
    $('#tombolTambahPermintaanPerbaikan').click(function(){
      $('#saveBtnPermintaanPerbaikan').addClass('show');
      $('#saveBtnUpdatePermintaanPerbaikan').attr('hidden', 'true'); 
      $('#saveBtnPermintaanPerbaikan').removeAttr('hidden'); 
    });


    // Tambah
    $('#saveBtnTambahPermintaanPerbaikan').click( function(e){
      
      e.preventDefault();
      $(this).html('Sending..');
  
      $.ajax({
        data: $('#formtambahpermintaanperbaikan').serialize(),
        url: "/dashboard/form/permintaan-perbaikan",
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
                    $('#saveBtnTambahPermintaanPerbaikan').html('Tambah');
                  });
                }else{
                  $('#formtambahpermintaanperbaikan').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahPermintaanPerbaikan').html('Tambah');
                  $('#tombolTambahPermintaanPerbaikan').trigger('click');
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses tambah data !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                    $(this).remove(); 
                 });
                    
                }
        },
        
      });

    })

   

    //edit
    $('body').on('click', '.editPermintaanPerbaikan', function () {
      $('#saveBtnTambahPermintaanPerbaikan').attr('hidden', 'true'); 
      $('#saveBtnUpdatePermintaanPerbaikan').removeAttr('hidden'); 
      $('#saveBtnUpdatePermintaanPerbaikan').html('Update');




      var permintaan_perbaikan_id = $(this).data('id');
      $.get("/dashboard/form/permintaan-perbaikan" +'/' + permintaan_perbaikan_id +'/edit', function (data) {
        $('#tambahpermintaanperbaikan').addClass('show');

          $('#id_permintaan_perbaikan').val(data.id);
          $('#jenis_mesin').val(data.jenis_mesin);
          $('#teknisi').val(data.teknisi_id);
          $('#user').val(data.user_id);
          $('#tipe_perbaikan').val(data.tipe_perbaikan);
          $('#masalah_perbaikan').val(data.masalah_perbaikan);
          $('#kronologis_kejadian').val(data.kronologis_kejadian);
          $('#analisa').val(data.analisa);
          $('#tindakan').val(data.tindakan);
          $('#progress').val(data.progress);
          $('#keterangan').val(data.keterangan);
          $('#mulai_terjadi').val(data.mulai_terjadi);
          $('#mulai_perbaikan').val(data.mulai_perbaikan);
          $('#selesai_perbaikan').val(data.selesai_perbaikan);
          $('#pemberi_order').val(data.pemberi_order);

          $("html, body").animate({ scrollTop: 0 }, "slow");
          // $('#scroll-ku').scrollTop = 0;
          
      })
   });


    //update
    $('#saveBtnUpdatePermintaanPerbaikan').click(function (e) {

      e.preventDefault();
      $(this).html('Sending..');
      var id_permintaan_perbaikan = $('#id_permintaan_perbaikan').val();
      $.ajax({
        data: $('#formtambahpermintaanperbaikan').serialize(),
        url: "/dashboard/form/permintaan-perbaikan"+'/'+id_permintaan_perbaikan,
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
                    $('#updateBtnUpdatePermintaanPerbaikan').html('Update');
                  });
                }else{
                  $(this).html('Update');
                  $('#formtambahpermintaanperbaikan').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahPermintaanPerbaikan').html('Tambah');
                  $('#tombolTambahPermintaanPerbaikan').trigger('click');   
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses Update !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                     $(this).remove(); 
                  });
                }
        },
        
      });

    })


    //batal
   $('#saveBtnBatalPermintaanPerbaikan').click( function(e){
     $(this).html('Sending..');
     e.preventDefault();
     $('#formtambahpermintaanperbaikan').trigger("reset");
        $('#saveBtnBatalPermintaanPerbaikan').html('Batal');
        $('#tombolTambahPermintaanPerbaikan').trigger('click');
        $(document).find('span.error-text').text('');

   })


    //delete
    $('body').on('click', '.deletePermintaanPerbaikan', function () {
     
        var permintaan_perbaikan_id = $(this).data("id");
        var r = confirm("Are You sure want to delete !");
         if (r == true) {
           $.ajax({
            type: "DELETE",
            url: "/dashboard/form/permintaan-perbaikan"+'/'+permintaan_perbaikan_id,
            success: function (data) {
              $('.table-responsive').before(`<div class="alert alert-danger alert-dismissible fade show" role="alert"> Sukses hapus!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
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