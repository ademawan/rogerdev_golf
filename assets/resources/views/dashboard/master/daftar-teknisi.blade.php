@extends('dashboard.layouts.main')

@section('container')
<div class="row">
  <div class="d-flex justify-content-center flex-wrap flex-md-nowrap align-items-center border-bottom">
    <h1 class="h2 d-blockr">DAFTAR PEKERJA</h1>
      {{-- <div class="btn-toolbar mb-2 mb-md-0">
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-secondary">Share</button>
          <button type="button" class="btn btn-sm btn-outline-secondary">Export</button>
        </div>
        <button type="button" class="btn btn-sm btn-outline-secondary dropdown-toggle">
          <span data-feather="calendar"></span>
          This week
        </button>
      </div> --}}
    </div>
  
    <div class="row my-3">
      <div class="col-md-12">
        <button type="button" id="tombolTambahTeknisi" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " data-bs-toggle="collapse" data-bs-target="#tambahteknisi" aria-expanded="false" aria-controls="collapseExample">
          Tambah Teknisi
        </button>
      </div>
        
      <div class="col-md-12 collapse" id="tambahteknisi">
        <div class="col-md">
          <form class="form-horizontal form-disable-enter" name="formtambahteknisi" id="formtambahteknisi">
            @csrf
            <input type="hidden" id="id_teknisi" name="id_teknisi" class="form-control" value="">

            <div class="row justify-content-around">
         
              <div class="col-lg-5 shadow-sm">
              
                <div data-spy="scroll" data-bs-target="#scroll-ku" data-bs-offset="0" style="height: 300px;overflow: auto;">
                  <div class="row py-2" style="margin:0 !important;">
                    <div class="col-md-12" >
                      <div class="mb-3 row">
                        <label for="inputtext" class="col-sm-3 col-form-label">Nama Lengkap</label>
                        <div class="col-sm-9">
                          <input type="text" name="nama_lengkap" class="form-control" id="nama_lengkap">
                          <span class="text-danger error-text nama_lengkap_error"></span>

                        </div>
                      </div>
                    </div>
                    <div class="col-md-12">
                      <div class="mb-3 row">
                        <label for="inputtext" class="col-sm-3 col-form-label">Jabatan</label>
                        <div class="col-sm-9">
                          <select id="jabatan" name="jabatan" class="form-select" aria-label="Default select example">
                            <option value="1">No Jabatan</option>
                            <option value="2">No Jabatan</option>
                            <option value="3">No Jabatan</option>
                          </select>
                          <span class="text-danger error-text jabatan_error"></span>

                        </div>
                      </div>      
                    </div>
                   
                    <div class="col-md-12">
                      <div class="mb-3 row">
                        <label for="inputtext" class="col-sm-3 col-form-label">Status Pekerja</label>
                        <div class="col-sm-9">
                          <select id="status" name="status" class="form-select" aria-label="Default select example">
                            <option value="1">Aktif</option>
                            <option value="0">Tidak Aktif</option>
                          </select>
                          <span class="text-danger error-text status_error"></span>

                        </div>
                      </div>      
                    </div>
  
                  </div>
                      <div class="d-grid gap-2 d-md-flex justify-content-md-center telusuri">
                        <button id="saveBtnUpdateTeknisi" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                        <button id="saveBtnTambahTeknisi" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                        <button id="saveBtnBatalTeknisi" class="btn btn-danger" type="submit">Batal</button>
                    </div>
         
                </div>
              </div>
              <div class="col-lg-6 shadow-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic maxime quidem modi repellat fugit velit explicabo nostrum sequi, blanditiis eum voluptas deserunt culpa nemo ut recusandae at nulla illo officiis dolores, cupiditate deleniti vel libero. Accusantium fugit officiis repellat eos sed, voluptatem facere reiciendis eaque inventore cupiditate repellendus veniam veritatis corrupti dignissimos voluptas. Eveniet obcaecati iste dignissimos autem.
                <div class="row">
                  <small>
                  <span class="text-danger error-text nama_lengkap_errornotif"></span>
                  <span class="text-danger error-text jabatan_errornotif"></span>
                  <span class="text-danger error-text status_errornotif"></span>
                  </small>
                </div>

              </div>
            </div>
  
            </div>
  
        </form>
  
        </div>
      </div>
  
    
    <div class="table-responsive">
      <table id="myTable" class="table table-striped w-100">
        <thead>
          <tr>
            <th scope="col">#</th>
            <th scope="col">Nama Lengkap</th>
            <th scope="col">Jabatan</th>
            <th scope="col">Status Pekerja</th>
            <th scope="col">Action</th>
          </tr>
        </thead>
        <tbody>
          
        </tbody>
      </table>
    </div>
      

</div>

@endsection

@push('script')
<script type="text/javascript" src="/js/master/teknisi.js"></script>
    
@endpush