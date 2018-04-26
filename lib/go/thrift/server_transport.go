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

// Server transport. Object which provides client transports.
// 网络层接口，通过Accept产生提供各种客户端transports
type TServerTransport interface {
	Listen() error
	Accept() (TTransport, error)
	Close() error

	// Optional method implementation. This signals to the server transport
	// that it should break out of any accept() or listen() that it is currently
	// blocked on. This method, if implemented, MUST be thread safe, as it may
	// be called from a different thread context than the other TServerTransport
	// methods.
	// 附加方法实现。用来通知server transport，需要停止accept()或listen()。如果实现了这个方法，必须保证是线程安全的，不同于其他transport方法，这个方法有可能从一个不同的线程上下文调用
	Interrupt() error
}
