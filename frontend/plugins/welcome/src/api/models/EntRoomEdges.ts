/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntRoomuse,
    EntRoomuseFromJSON,
    EntRoomuseFromJSONTyped,
    EntRoomuseToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoomEdges
 */
export interface EntRoomEdges {
    /**
     * Roomuses holds the value of the roomuses edge.
     * @type {Array<EntRoomuse>}
     * @memberof EntRoomEdges
     */
    roomuses?: Array<EntRoomuse>;
}

export function EntRoomEdgesFromJSON(json: any): EntRoomEdges {
    return EntRoomEdgesFromJSONTyped(json, false);
}

export function EntRoomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomuses': !exists(json, 'roomuses') ? undefined : ((json['roomuses'] as Array<any>).map(EntRoomuseFromJSON)),
    };
}

export function EntRoomEdgesToJSON(value?: EntRoomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomuses': value.roomuses === undefined ? undefined : ((value.roomuses as Array<any>).map(EntRoomuseToJSON)),
    };
}


