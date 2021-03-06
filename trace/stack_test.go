/*
 * Licensed to the OpenSkywalking under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenSkywalking licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package trace

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	assert := assert.New(t)
	s :=new(stack)
	s.push("1")
	s.push("2")
	s.push("3")

	assert.Equal("1",s.getLast())
	assert.Equal("3",s.pop())
	assert.Equal("2",s.pop())
	assert.Equal("1",s.pop())
}

func Test_A(T *testing.T){
	//conn, err := grpc.Dial("", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := pb.NewInstanceDiscoveryServiceClient(conn)
	//
	//// Contact the server and print out its response.
	//name := ""
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//c
	//r, err := c.Heartbeat(ctx,nil,)
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//
	//r.
	//time.Sleep(time.Second)
	//var a chan  int
	//close(a)
}