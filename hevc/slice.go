/* The copyright in this software is being made available under the BSD
 * License, included below. This software may be subject to other third party
 * and contributor rights, including patent rights, and no such rights are
 * granted under this license.
 *
 * Copyright (c) 2012-2013, H265.net
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *  * Neither the name of the H265.net nor the names of its contributors may
 *    be used to endorse or promote products derived from this software without
 *    specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS
 * BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF
 * THE POSSIBILITY OF SUCH DAMAGE.
 */
 
package hevc

import (
    //"fmt"
    "errors"
    "bufio"
    "container/list"
)

type Slice struct {
    ParameterSet
    M_frame                 *Frame
    PicParameterSetId       int
    FirstSliceInPicFlag     int
    SliceAddr				int
}

func NewSlice(reader *bufio.Reader) *Slice{
	var slice Slice
	
	slice.ParameterSet.Reader = reader
	slice.ParameterSet.FieldList = list.New()
	
    return &slice
}

func (ps *Slice) Parse() (line string, err error)  {
    line, err = ps.ParameterSet.Parse()

    ps.PicParameterSetId = -1
    ps.FirstSliceInPicFlag = 0
    for e := ps.FieldList.Front(); e != nil; e = e.Next() {
        field := e.Value.(Field)

        if field.Name == "pic_parameter_set_id" {
            ps.PicParameterSetId = field.Value
        }else if field.Name == "first_slice_in_pic_flag" {
            ps.FirstSliceInPicFlag =  field.Value
        }
    }

    if ps.PicParameterSetId < 0 {
        err = errors.New("pic_parameter_set_id not found or invalid")
    }
    
    if ps.FirstSliceInPicFlag==1{
    	ps.SliceAddr = 0
    }else{
    	ps.SliceAddr, err = ps.ParameterSet.GetValue("slice_address") 
    }

    return line, err
}
