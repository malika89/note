from faker import Faker
import unittest
from faker.providers import BaseProvider
import faker_microservice
import pymysql

# 自定义项目名(项目名格式x-y)
class MyProvider(BaseProvider):
    def app_name(self):
        return f"{fake.pystr(6)}-{fake.pystr(5)}"


fake = Faker()
# then add new provider to faker instance
fake.add_provider(MyProvider)
fake.add_provider(faker_microservice.Provider)


class Create_data(object):

    def make_sigle_data(self):
        """
        cpu,app,pid,path,port,name(主机名),ip
        :return:
        """
        app_name = fake.microservice()
        data_dict = {
            "name": fake.pystr(),
            "ip": fake.ipv4(),
            "cpu": fake.random_int(min=1,max=32),
            "app": app_name,
            "pid": fake.random_int(max=1000),
            "port": fake.random_int(min=1,max=65535),
            "path": f"/web/apps/{app_name}",

        }
        return data_dict

    def get_dataN(self, n):
        data = [self.make_sigle_data() for i in range(n)]
        return data

    def convert_sql(self, table_name, data_dict=dict()):
        field_columns = ",".join('`{}`'.format(k) for k in data_dict.keys())
        val_columns = ",".join('%({})s'.format(k) for k in data_dict.keys())
        coverted_sql = f"insert into {table_name}({field_columns}) values({val_columns})"
        return coverted_sql

    def write_batch(self, db_dict, n):
        db = pymysql.connect(**db_dict)
        cursor = db.cursor()
        data_lst = self.get_dataN(n)
        # 每次2000条写入：
        if n<= 2000:
            cursor.executemany(self.convert_sql(data_lst[0]), data_lst)
        else:
            tmp = n/2000
            for i in range(0, tmp):
                start = i*2000
                end = (i+1)*2000 if (i+1)*2000<n else n
                cursor.executemany(self.convert_sql(data_lst[0]), data_lst[start:end])
        db.commit()
        cursor.close()
        db.close()


class Test_faker(unittest.TestCase):
    def test_make(self):
        self.assertTrue(Create_data().make_sigle_data())


if __name__ == '__main__':
    # unittest.main
    db_config = {"host": "127.0.0.1", "port": 3306, "user": "root", "password": "case123"}
    Create_data().write_batch(db_config, 10000)

