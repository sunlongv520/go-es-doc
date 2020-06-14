<template>
    <div>
      <div>
          <dl class="sdt">
              <dt>
                  出版社
              </dt>
              <dd v-for="item in pressList"><a href="">{{item}}</a></dd>
          </dl>
      </div>
        <div>
            <el-table
                    :data="bookList"
                    style="width: 100%">
                <el-table-column
                        prop="BookId"
                        label="ID"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="BookName"
                        label="商品名称"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="price1"
                        label="价格">
                </el-table-column>
                <el-table-column
                        prop="KindName"
                        label="分类">
                </el-table-column>
            </el-table>
        </div>

    </div>

</template>
<script>
   module.exports ={
       data(){
           return {
               bookList: [],
               pressList: []
           }
       },
       created(){
            this.loadPress()  // 加载出版社列表
       },
       methods: {
           async loadPress(){
               const response = await axios.get("/helper/press")
               const { result } = response.data
               this.pressList = result
           }
       }
   }
</script>
<style>
  .sdt{margin: 0 auto ;width:90%;border-radius: 5px;}
  .sdt dt{width:100%;display: block;color:#3A7B43;font-size:16px;font-weight: bold}
  .sdt dd{float:left;margin:0 auto;text-indent:1em}
  .sdt dd a{color: #3a8ee6}
  .sdt dd a:hover{background: #eb5975;color:#fff}
</style>