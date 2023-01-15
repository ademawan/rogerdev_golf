@extends('dashboard.layouts.main')

@section('container')
<div class="row" id="container">
<div class="d-flex justify-content-center flex-wrap flex-md-nowrap align-items-center border-bottom">
    <h1 class="h2 d-blockr">DAFTAR PERMINTAAN PERBAIKAN</h1>

  </div>
  <div class="row my-3">
    <div class="col-md-12">
      <button type="button" id="tombolTambahPermintaanPerbaikan" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " data-bs-toggle="collapse" data-bs-target="#tambahpermintaanperbaikan" aria-expanded="false" aria-controls="collapseExample">
        Tambah Permintaan Perbaikan
      </button>
      <div hidden>
        <button type="button" id="tombolRefreshTable" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " >
          Refresh Table
        </button>

      </div>
    </div>
      
    <div class="col-md-12 collapse" id="tambahpermintaanperbaikan">
      <div class="col-md">
        <form class="form-horizontal form-disable-enter" name="formtambahpermintaanperbaikan" id="formtambahpermintaanperbaikan">
          @csrf
          <input type="hidden" id="id_permintaan_perbaikan" name="id_permintaan_perbaikan" class="form-control" value="">
          <input type="hidden" id="id_user" name="id_user" class="form-control" value="{{ auth()->user()->id }}">

          <div class="row justify-content-around">
       
            <div class="col-lg-5 shadow-sm">
            
              <div data-spy="scroll" id="scroll-ku" data-bs-target="#scroll-ku" data-bs-offset="0" style="height: 300px;overflow: auto;">
                <div class="row py-2" style="margin:0 !important;">

                  <div class="col-md-12" >
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Jenis Mesin</label>
                      <div class="col-sm-9">
                        <div class="input-group">
                          <select type="text" data-width="100%" id="selectjenismesin" name="jenis_mesin" class="form-control" autocomplete="off">
                          
                            <option value="">Pilih Mesin</option>
                            @foreach ($mesins as $mesin)
                                
                            <option value="{{ $mesin->id }}">{{ $mesin->jenis_mesin }}</option>
                            @endforeach
                            
                          </select>
                          <span class="text-danger error-text jenis_mesin_error"></span>
  
                        </div>
                        </div>
                    </div>
                  </div>

                  <div class="col-md-12" >
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Teknisi</label>
                      <div class="col-sm-9">
                        <div class="input-group">
                          <select type="text" data-width="100%" id="selectteknisi" name="teknisi" class="form-control" autocomplete="off">
                          
                            <option value="">Pilih Teknisi</option>
                            @foreach ($teknisis as $teknisi)
                                
                            <option value="{{ $teknisi->id }}">{{ $teknisi->nama_lengkap }}</option>
                            @endforeach
                            
                          </select>
                          <span class="text-danger error-text teknisi_error"></span>
  
                        </div>
                        </div>
                    </div>
                  </div>
                  
                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Tipe Perbaikan</label>
                      <div class="col-sm-9">
                        <input type="text" id="tipe_perbaikan" name="tipe_perbaikan" class="form-control">
                         <span class="text-danger error-text tipe_perbaikan_error"></span>
                      </div>
                    </div>
                  </div>
                  
                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Masalah Kerusakan</label>
                      <div class="col-sm-9">
                        <input type="text" id="masalah_kerusakan" name="masalah_kerusakan" class="form-control">
                         <span class="text-danger error-text masalah_kerusakan_error"></span>
                      </div>
                    </div>
                  </div>
                  

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Kronologis Kejadian</label>
                      <div class="col-sm-9">
                        <input type="text" id="kronologis_kejadian" name="kronologis_kejadian" class="form-control">
                          <span class="text-danger error-text kronologis_kejadian_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Analisa</label>
                      <div class="col-sm-9">
                        <input type="text" id="analisa" name="analisa" class="form-control">
                          <span class="text-danger error-text analisa_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Tindakan</label>
                      <div class="col-sm-9">
                        <input type="text" id="tindakan" name="tindakan" class="form-control">
                          <span class="text-danger error-text tindakan_error"></span>
                      </div>
                    </div>
                  </div>


                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Progress</label>
                      <div class="col-sm-9">
                        <input type="text" id="progress" name="progress" class="form-control">
                          <span class="text-danger error-text progress_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Keterangan</label>
                      <div class="col-sm-9">
                        <textarea class="form-control" id="keterangan" name="keterangan" id="floatingTextarea2" style="height: 70px" ></textarea>
                        <span class="text-danger error-text keterangan_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Mulai Terjadi</label>
                      <div class="col-sm-9">
                        <input type="date" id="mulai_terjadi" name="mulai_terjadi" class="form-control">
                          <span class="text-danger error-text mulai_terjadi_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Mulai Perbaikan</label>
                      <div class="col-sm-9">
                        <input type="date" id="mulai_perbaikan" name="mulai_perbaikan" class="form-control">
                          <span class="text-danger error-text mulai_perbaikan_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Selesai Perbaikan</label>
                      <div class="col-sm-9">
                        <input type="date" id="selesai_perbaikan" name="selesai_perbaikan" class="form-control">
                          <span class="text-danger error-text selesai_perbaikan_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Pemberi Order</label>
                      <div class="col-sm-9">
                        <input type="text" id="pemberi_order" name="pemberi_order" class="form-control">
                          <span class="text-danger error-text pemberi_order_error"></span>
                      </div>
                    </div>
                  </div>
                  
                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Sparepart ID</label>
                      <div class="col-sm-4 px-0" >
                        <input type="text" id="sparepart" name="sparepart[]" class="form-control inputsparepart" readonly>
                          <span class="text-danger error-text sparepart_error"></span>
                      </div>
                      <div class="col-sm-3 px-1">
                        <input type="text" id="qty" name="qty[]" class="form-control" placeholder="QTY" style="display:block !important; padding:5px !important;" readonly>
                          <span class="text-danger error-text qty_error"></span>
                      </div>
                      <div class="col-sm-1 px-0">
                        <button type="button" class="buttonCari" style="display:block !important; padding:5px !important;" data-bs-toggle="modal" data-bs-target="#modalSparepart">
                          Cari
                        </button>
                      </div>
                      <div class="col-sm-1 px-0">
                        <button type="button" class=" d-block padding-2 buttonTambahSparepart" style="display:block !important; padding:5px !important;">
                          +
                        </button>
                      </div>
                    </div>
                  </div>
                  
                  
                </div>
                <div class="d-grid gap-2 d-md-flex justify-content-md-center telusuri">
                    <button id="saveBtnUpdatePermintaanPerbaikan" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                    <button id="saveBtnTambahPermintaanPerbaikan" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                    <button id="saveBtnBatalPermintaanPerbaikan" class="btn btn-danger updateBtnBatalPermintaanPerbaikan" type="submit">Batal</button>
                </div>
       
              </div>
            </div>
            <div class="col-lg-6 shadow-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic maxime quidem modi repellat fugit velit explicabo nostrum sequi, blanditiis eum voluptas deserunt culpa nemo ut recusandae at nulla illo officiis dolores, cupiditate deleniti vel libero. Accusantium fugit officiis repellat eos sed, voluptatem facere reiciendis eaque inventore cupiditate repellendus veniam veritatis corrupti dignissimos voluptas. Eveniet obcaecati iste dignissimos autem.
              <div class="row">
                <small>
                <span class="text-danger error-text jenis_mesin_errornotif"></span>
                <span class="text-danger error-text teknisi_errornotif"></span>
                <span class="text-danger error-text user_errornotif"></span>
                <span class="text-danger error-text tipe_perbaikan_errornotif"></span>
                <span class="text-danger error-text masalah_kerusakan_errornotif"></span>
                <span class="text-danger error-text kronologis_kejadian_errornotif"></span>
                <span class="text-danger error-text analisa_errornotif"></span>
                <span class="text-danger error-text tindakan_errornotif"></span>
                <span class="text-danger error-text progress_errornotif"></span>
                <span class="text-danger error-text keterangan_errornotif"></span>
                <span class="text-danger error-text mulai_terjadi_errornotif"></span>
                <span class="text-danger error-text mulai_perbaikan_errornotif"></span>
                <span class="text-danger error-text selesai_perbaikan_errornotif"></span>
                <span class="text-danger error-text pemberi_order_errornotif"></span>
                </small>
              </div>
            </div>
            
          </div>

          </div>

      </form>

      </div>
    </div>
    
  </div>
  <hr class="dropdown-divider mb-3">
 
    <div class="table-responsive">
      <table id="myTable" class="table table-striped" width="100%">
        <thead>
          <tr>
            <th>Nama Mesin</th>
            <th>Nama Teknisi</th>
            <th>Nama User</th>  
            <th>Tipe Perbaikan</th>  
            <th>Masalah Kerusakan</th>  
            <th>Kronologis Kejadian</th>  
            <th>Analisa</th>  
            <th>Tindakan</th>  
            <th>Progress</th>            
            <th>Keterangan</th>
            <th>Mulai Terjadi</th>
            <th>Mulai Perbaikan</th>
            <th>Selesai Perbaikan</th>
            <th>Pemberi Order</th>
          </tr>     
        </thead>
        <tbody>
          
        </tbody>
      </table>
    </div>


    {{-- modal --}}
    <!-- Button trigger modal -->


<!-- Modal -->
<div class="modal fade" id="modalSparepart" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLabel">Pilih Sparepart</h5>
      </div>
      <div class="modal-body">
        
        <div class="col-md-12" id="sparepart2" >
          <div class="mb-3 row">
            <label for="inputtext" class="col-sm-3 col-form-label">Sparepart</label>
            <div class="col-sm-5">
              <div class="input-group">
                <select type="text" data-width="80%" id="selectsparepart" name="selectsparepart[]" class="form-select" autocomplete="off">
                
                  <option value="">Pilih Sparepart</option>
                  @foreach ($spareparts as $sparepart)
                      
                  <option class="optionValue" value="{{ $sparepart->id }}">{{ $sparepart->nama }}</option>
                  @endforeach
                  
                </select>
                <span class="text-danger error-text selectsparepart_error"></span>

              </div>
              </div>
              <div class="col-sm-2">
                <input type="text" id="qty" name="qty[]" class="form-control inputQty" placeholder="QTY">

              </div>
          </div>
        </div>
      </div>
      <div class="modal-footer">
        <button type="button" id="buttonModalBatal" class="btn btn-secondary" data-bs-dismiss="modal">Batal</button>
        <button type="button" id="buttonModalTambah" class="btn btn-primary">Tambah</button>
      </div>
    </div>
  </div>
</div>

</div>    
@endsection

@push('script')
<script src="https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/js/select2.min.js"></script>

<script src="/js/form/permintaan_perbaikan.js"></script>

    
@endpush