class StatExporterProcess(Process):

    def __init__(self,endpoint,stat_store_spans):
        super(StatExporterProcess,self).__init__()
        self.testDataStore = MongoClient('mongodb://dongxincbms:bms2016!@127.0.0.1/admin')
        self.endpoint = endpoint
        self.statExporter = StatExporter(stat_store_spans)

    def run(self):
        while True:
            for rec in self.testDataStore['tmp'][self.endpoint].find({}):
                app_name = rec['app_name']
                intf_name = rec['intf_name']
                _id = rec['_id']
                rec['_id'] = rec['_id_']
                del rec['app_name']
                del rec['intf_name']
                self.statExporter.export(app_name, intf_name, [rec])
                self.testDataStore['tmp'][self.endpoint].remove(spec_or_id={'_id': ObjectId(_id)},safe=True)
            time.sleep(0.001)
