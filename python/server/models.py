from conf.base import BaseDB,engine
import sys

from sqlalchemy import(
    Column,
    Integer,
    String,
    DateTime
)

class Users(BaseDB):
    """ table for Users
    """
    __tablename__ = "users"

    #定义表结构，包括id，phone，password，createTime
    id = Column(Integer,primary_key = True)
    phone = Column(String(14),nullable = False)
    password = Column(String(50),nullable = True)
    createTime = Column(DateTime,nullable = True)

    def __init__(self,phone,password,createTime):
        self.phone = phone
        self.password = password
        self.createTime = createTime

def initdb():
    BaseDB.metadata.create_all(engine)

if __name__ == "__main__":
    print ("Initialize database")
    initdb()