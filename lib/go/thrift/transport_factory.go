/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package thrift

// Factory class used to create wrapped instance of Transports.
// This is used primarily in servers, which get Transports from
// a ServerTransport and then may want to mutate them (i.e. create
// a BufferedTransport from the underlying base transport)
// Fatctory类用于创建Trasnports的复写实例
// 主要用在servers里，从ServerTransport获取Transports，然后更变他们(比如，基于基础transport，创建BufferedTransport)
type TTransportFactory interface {
	GetTransport(trans TTransport) (TTransport, error)
}

type tTransportFactory struct{}

// Return a wrapped instance of the base Transport.
// 返回基础Transport的复写实例
func (p *tTransportFactory) GetTransport(trans TTransport) (TTransport, error) {
	return trans, nil
}

func NewTTransportFactory() TTransportFactory {
	return &tTransportFactory{}
}
